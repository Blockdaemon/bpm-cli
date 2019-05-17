package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"io/ioutil"
	"os"
	"bufio"
)

// ----------------------------------------------------------------------------
// BasicManager
// ----------------------------------------------------------------------------

type BasicManager struct {
	cli *client.Client
}

func NewBasicManager() (*BasicManager, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return &BasicManager{
		cli: cli,
	}, nil
}

// RemoveContainer stops and removes a container
func (bm *BasicManager) ContainerAbsent(ctx context.Context, containerName string) error {
	running, err := bm.isContainerRunning(ctx, containerName)
	if err != nil {
		return err
	}

	if running {
		fmt.Printf("Stopping container '%s'\n", containerName)

		if err := bm.cli.ContainerStop(ctx, containerName, nil); err != nil {
			return err
		}
	}

	exists, err := bm.doesContainerExists(ctx, containerName)
	if err != nil {
		return err
	}

	if exists {
		fmt.Printf("Removing container '%s'\n", containerName)
	
		if err := bm.cli.ContainerRemove(ctx, containerName, types.ContainerRemoveOptions{ RemoveVolumes: true, }); err != nil {
			return err
		}
	}

	return nil
}

func (bm *BasicManager) NetworkExists(ctx context.Context, networkID string) error {
	_, err := bm.cli.NetworkInspect(ctx, networkID)

	if client.IsErrNetworkNotFound(err) { 
		_, err := bm.cli.NetworkCreate(ctx, networkID, types.NetworkCreate{ CheckDuplicate: true, })

		if err != nil {
			return err
		}
	}

	return err
}

type ContainerMountParameter struct {
	From string;
	To string;
}

type ContainerPortParameter struct {
	HostIP string;
	HostPort string;
	ContainerPort string;
	Protocol string;
}



type ContainerRunParameters struct {
	ContainerName string;
	ContainerImage string;
	NetworkName string;
	EnvFilename string;
	Mounts []ContainerMountParameter;
	Ports []ContainerPortParameter;

}

func (bm *BasicManager) doesContainerExists(ctx context.Context, containerName string) (bool, error) {
	_, err := bm.cli.ContainerInspect(ctx, containerName)
	if err != nil {
		if client.IsErrContainerNotFound(err) { 
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (bm *BasicManager) isContainerRunning(ctx context.Context, containerName string) (bool, error) {
	inspect, err := bm.cli.ContainerInspect(ctx, containerName)
	if err != nil {
		if client.IsErrContainerNotFound(err) { 
			return false, nil // a non existing container is not running!
		}

		return false, err
	}

	return inspect.State.Running, nil

}

// ContainerRuns starts a container if it doesn't run yet
func (bm *BasicManager) ContainerRuns(ctx context.Context, params ContainerRunParameters) error {
	// Pull an image from the repository
	out, err := bm.cli.ImagePull(ctx, params.ContainerImage, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer out.Close()
    if _, err := ioutil.ReadAll(out); err != nil {
            panic(err)
    }

	exists, err := bm.doesContainerExists(ctx, params.ContainerName)
	if err != nil {
		return err
	}
	if exists {
		fmt.Printf("Container '%s' already exists. Skipping creation.\n", params.ContainerName)
	} else {
		fmt.Printf("Creating container '%s'\n", params.ContainerName)	

		// Environment variable configuration
		var envs []string
		if params.EnvFilename != "" {
			envs, err = readLines(params.EnvFilename)
			if err != nil {
				return err
			}
		}

		// Container's config
		containerCfg := &container.Config{
			Image: params.ContainerImage,
			Env: envs,
		}

		// Host's config

		portBindings := make(map[nat.Port][]nat.PortBinding)


		for _, portParameter := range params.Ports {
			containerPort, err := nat.NewPort(portParameter.Protocol, portParameter.ContainerPort)
			if err != nil {
				return err
			}

			portBindings[containerPort] = []nat.PortBinding{
				nat.PortBinding{
					HostIP: portParameter.HostIP,
					HostPort: portParameter.HostPort,
				},
			}

		}


		hostCfg := &container.HostConfig{
			PortBindings: portBindings,

			RestartPolicy: container.RestartPolicy{
				Name: "unless-stopped",
			},

			LogConfig: container.LogConfig{
				Type: "json-file",
				Config: map[string]string{
					"max-size": "10m",
					"max-file": "3",
				},
			},
		}

		// Mountpoints
		for _, mountParam := range params.Mounts {
			hostCfg.Mounts = append(hostCfg.Mounts, mount.Mount {
                Type:   mount.TypeBind,
                Source: mountParam.From,
                Target: mountParam.To,
			})
		}

		// Network configuration
		endpointsConfig := make(map[string]*network.EndpointSettings)
		endpointsConfig[params.NetworkName] = &network.EndpointSettings{
			NetworkID: params.NetworkName,
		}
		networkConfig := &network.NetworkingConfig{
			EndpointsConfig: endpointsConfig,
		}

		// Create a container with configs
		_, err := bm.cli.ContainerCreate(ctx, containerCfg, hostCfg, networkConfig, params.ContainerName)
		if err != nil {
			return err
		}
	}

	running, err := bm.isContainerRunning(ctx, params.ContainerName) 
	if err != nil {
		return err
	}
	if running {
		fmt.Printf("Container '%s' is already running. Skipping start.\n", params.ContainerName)
	} else {
		fmt.Printf("Starting container '%s'\n", params.ContainerName)	

		if err := bm.cli.ContainerStart(ctx, params.ContainerName, types.ContainerStartOptions{}); err != nil {
			return err
		}
	}

	return nil
}

func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}


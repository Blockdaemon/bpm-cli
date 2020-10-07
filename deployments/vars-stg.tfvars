gcp_project                = "blockdaemon-staging"
gcp_bucket_name            = "stg-bpm-cli-docs"
k8s_context                = "gke_blockdaemon-staging_us-west2-a_app-cluster-stg"
k8s_namespace              = "bpm-cli-docs"
k8s_image_pull_secret_name = "regcred"

nginx_ingress_loadbalancer = "stg.api.blockdaemon.com"

dns_env_subdomain = "stg."
dns_domain = "cli.bpm.docs.blockdaemon.com"

redoc_image = "redocly/redoc:v2.0.0-rc.41"


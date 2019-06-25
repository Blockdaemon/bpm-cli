




# Code structure

For easier testability we separate business logic from Cobra commands.

- `cmd/bpm/main.go` is the main entrypint
- `internal/bpm/cmd/` contains the Cobra commands. These are only responsible for:
	- Parsing arguments
	- Calling a function in `internal/bpm/tasks`
	- Returning an error or printing the output
- `internal/bpm/tasks/` contains the business logic. Each function performs a single task and return either an error or a string containing the output of the task. This makes it very simple to test the tasks.
- `pkg/models/` contains low-level code and models. This is typically called by the tasks.

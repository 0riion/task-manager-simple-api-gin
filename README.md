## Project structure

This is a project structure for a Go/Gin API for managing tasks. Here is a simplified breakdown of the project structure:

```bash
@julio ➜ task-manager-api-golang git(main) tree
.
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── Makefile
├── pkg
│   ├── common
│   │   ├── config
│   │   │   └── config.go
│   │   ├── db
│   │   │   └── db.go
│   │   ├── envs
│   │   └── models
│   │       └── task.go
│   └── tasks
│       ├── controller.go
│       ├── create_task.go
│       ├── delete_task.go
│       ├── list_tasks.go
│       ├── retrieve_task.go
│       └── update_task.go
└── README.md

8 directories, 14 files
@julio ➜ task-manager-api-
```

- `cmd`: This folder contains the main file of the application that starts the server.

- `go.mod` and `go.sum`: These files define the project's dependencies and their versions.

- `Makefile`: This file contains a set of rules to build, test, and run the application.

- `pkg`: This folder contains the reusable packages that make up the application:

- `common`: This package contains common functionality used throughout the application, such as configuration, database access, environment variables, and data models.

- `tasks`: This package contains the code for managing tasks, including a controller and six sub-packages for creating, retrieving, updating, and deleting tasks.

Overall, the project is organized in a modular and easy-to-follow structure. The Makefile provides a convenient way to build and run the application, while the pkg folder separates the code into reusable packages.

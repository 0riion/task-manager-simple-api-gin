## Project Overview

This project is a Go/Gin API for managing tasks, with endpoints for creating, retrieving, updating, and deleting tasks. It provides a solid foundation for building more features and functionality upon.

## Project Details

The API is built with Go and Gin, and it allows users to perform CRUD (Create, Read, Update, and Delete) operations on tasks. It provides endpoints for creating new tasks, retrieving a single task or a list of tasks, updating a task, and deleting a task. The project is organized in a modular and easy-to-follow structure, making it easy to add more endpoints or functionality as needed.

This API can be used as a starting point for building a more complex task management system or integrated with other applications. With its simple and well-organized codebase, developers can easily extend its functionality to meet their specific needs.

## Project structure

This is a project structure for a Go/Gin API for managing tasks. Here is a simplified breakdown of the project structure:

```bash
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
```

- `cmd`: This folder contains the main file of the application that starts the server.

- `go.mod` and `go.sum`: These files define the project's dependencies and their versions.

- `Makefile`: This file contains a set of rules to build, test, and run the application.

- `pkg`: This folder contains the reusable packages that make up the application:

- `common`: This package contains common functionality used throughout the application, such as configuration, database access, environment variables, and data models.

- `tasks`: This package contains the code for managing tasks, including a controller and six sub-packages for creating, retrieving, updating, and deleting tasks.

Overall, the project is organized in a modular and easy-to-follow structure. The Makefile provides a convenient way to build and run the application, while the pkg folder separates the code into reusable packages.

## How to run the application

To run the application, you need to have Go installed on your machine. If you don't have it installed, you can download and install it from the official website.

After installing Go, you can clone the repository and start the application by running the following commands:

```bash
git clone https://github.com/0riion/task-manager-simple-api-gin.git
cd task-manager-simple-api-gin
make run
```

The `make run` command will build the application and start the server on port 8080. You can then access the API at http://localhost:8000.

## Implemented Programming Fundamentals

This project implements several core Go programming concepts, such as variables, functions, modules, packages, services, interfaces, structs, pointers, slices, maps, and more.

# Task Manager Go-Cli-App

A powerful Task Management application written in Go. This project provides both a **Command Line Interface (CLI)** and a **REST API** with integrated **Swagger UI**.

## Features

- **CLI Support**: Add, list, search, and update tasks from your terminal.
- **REST API**: Fully functional API built with Gin.
- **Swagger Documentation**: Interactive API testing via Swagger UI.
- **Persistence**: Data is saved locally in `tasks.json`.
- **Exporting**: Export your tasks to CSV files.

## Getting Started

### 1. Installation

Make sure you have Go installed, then clone the repository and install dependencies:

```bash
go mod tidy
```

### 2. Running the API (Serve Mode)

To start the REST API server on port 8080:

```bash
go run main.go serve
```

*   **API URL:** `http://localhost:8080/api/v1`
*   **Swagger UI:** `http://localhost:8080/swagger/index.html`

### 3. Using the CLI

To list your tasks:
```bash
go run main.go list
```

To add a new task:
```bash
go run main.go add --title "My First Task" --desc "Description here"
```

## Branch Information

You are currently on the `cli-code1` branch, which contains the latest REST API features.

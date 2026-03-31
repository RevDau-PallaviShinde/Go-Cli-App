# Task Manager CLI

A Go-based CLI application to manage daily tasks.

## Features

- Add Task
- List Task
- Update Task
- Delete Task
- Search Task
- Export CSV
- Summary
- Priority
- Due Date

## Installation

git clone https://github.com/RevDau-PallaviShinde/Go-Cli-App

cd Go-Cli-App

go mod tidy

go run main.go

## Commands

1. go run main.go add --title "Test"
2. go run main.go list
3. go run main.go list --status completed
4. go run main.go list --priority medium
5. go run main.go delete
6. go run main.go update --id 1 --status completed
7. go run main.go update --id 1 --priority high
8. go run main.go my-task-summary
9. go run main.go pallavi-search --keyword go

Add task

go run main.go add --title "Learn Go"

List tasks

go run main.go list

Update task

go run main.go update --id 1 --status completed

Delete task

go run main.go delete --id 1

Search

go run main.go pallavi-search --keyword go

Export

go run main.go export --file tasks.csv

Summary

go run main.go my-task-summary












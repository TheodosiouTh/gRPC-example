# gRPC-example

A simple todo application used for tracking tasks made using Golang, Cobra (for the CLI client) & PostgreSQL.

### What I wanted to learn

- How to connect a DB to a server made using Golang
- How gRPC works
- How to containerize Database and Server using Docker

## Requrements

- [Docker](https://docs.docker.com/): Used for running the server and database inside containers.

---

## Running the Project

1. Open the Terminal
2. Navigate to the Project's directory
3. Start the server & database inside a container:
   ```console
   $ make start
   ```
4. Open a second Terminal
5. Navigate to the Project's directory
6. Install the client:
   ```console
   $ go install
   ```

---

## Commands

- `todo list`: Lists all (non-removed) tasks.
- `todo add <task-name>`: Adds a task to the task list
- `todo check <task-id>`: Changes the status of a task to done
- `todo remove <task-id>`: Removes a task from the task list

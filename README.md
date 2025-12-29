# stack-gen-TODO

A minimal REST API for managing TODO items built using **Go (standard library only)**.  
This project demonstrates in-memory data storage, basic request validation, and unit testing with a strong focus on correctness and code clarity.

---

## Overview

The application exposes a small set of REST endpoints to **create, list, and delete TODO items**.  
All data is stored in memory, keeping the service simple and easy to reason about while emphasizing clean architecture and separation of concerns.

---

## Features

- Create TODO items
- Retrieve all TODO items
- Delete TODO items by ID
- In-memory storage (no database)
- Clean separation between handlers, services, and models
- Basic request validation
- Unit tests for business logic
- Uses **only Go’s standard library**

---

## Tech Stack

| Component | Technology |
|---------|------------|
| Language | Go |
| HTTP Server | `net/http` |
| JSON Handling | `encoding/json` |
| Testing | `testing` (standard library) |

---

## Setup & Run Instructions

### Prerequisites

| Requirement | Version |
|-----------|---------|
| Go | 1.20 or later |

Verify Go installation:
go version

## Running the Server

From the project root directory:

go run main.go

The server will start at:

http://localhost:8080

## API Endpoints

| Method | Endpoint        | Description             |
|------|-----------------|-------------------------|
| POST | /todos          | Create a new TODO       |
| GET  | /todos          | Retrieve all TODOs      |
| DELETE | /todos/{id}   | Delete a TODO by ID     |

## Running Tests
To run all unit tests:
go test ./...

### Example Requests

#### Create a TODO
```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"id":"1","title":"Learn Go","completed":false}'
```


Get All TODOs

```bash
curl http://localhost:8080/todos
```

Delete a TODO
```bash
curl -X DELETE http://localhost:8080/todos/1
```
Deletes the TODO with ID `1` (the last path parameter represents the TODO ID).

## Assumptions & Limitations

- Data is stored in memory and is lost when the server restarts
- No authentication or authorization is implemented
- No persistent storage or database
- Validation is minimal (ID and title must be non-empty)
- The Entire project is under the folder titled stack-gen-TODO 


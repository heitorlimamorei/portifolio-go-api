# Job Opening CRUD API

This is a Job Opening CRUD (Create, Read, Update, Delete) API implemented in Go lang using the Gin-Gonic web framework, SQLite as the database, and GORM as the Object Relational Mapping (ORM) library.

## Features

- **Create:** Add new job openings with details like title, description, location, salary, remote status, link, and company.
- **Read:** Retrieve information about existing job openings.
- **Update:** Modify the details of a specific job opening.
- **Delete:** Remove a job opening from the database.

## Technologies Used

- Go lang: The programming language used for implementing the API.
- Gin-Gonic: A web framework written in Go for building APIs.
- SQLite: A lightweight and self-contained SQL database engine.
- GORM: An ORM library for Go, providing a convenient way to interact with databases.

## Prerequisites

- [Go](https://golang.org/dl/): Make sure Go is installed on your machine.
- [Gin-Gonic](https://github.com/gin-gonic/gin): Install the Gin-Gonic framework.
- [GORM](https://gorm.io/docs/): Install the GORM library.
- [SQLite](https://www.sqlite.org/download.html): Install SQLite database.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/heitorlimamorei/portifolio-go-api.git
    cd job-opening-api
    ```

2. Install dependencies:

    ```bash
    go mod download
    ```

3. Run the API:

    ```bash
    go run main.go
    ```

The API will be running at `http://localhost:8080`.

## API Endpoints

### Create a Job Opening

```http
POST /opening
{
    "role": "Software Engineer",
    "location": "San Francisco, CA",
    "salary": 100000,
    "remote": true,
    "link": "https://example.com/job/software-engineer",
    "company": "TechCo"
}
```

### Get All Job Openings

```http
GET /openings
```

### Get Job Opening by ID

```http
GET /opening/:id
```

### Update Job Opening

```http
PUT /opening/:id
{
    "role": "Software Engineer",
    "location": "San Francisco, CA",
    "salary": 100000,
    "remote": true,
    "link": "https://example.com/job/software-engineer",
    "company": "TechCo"
}
```

### Delete Job Opening

```http
DELETE /opening/:id
```
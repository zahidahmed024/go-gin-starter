# Go Gin Starter

This is a boilerplate for a Go Gin application with JWT authentication, PostgreSQL, and Docker.

## Features

- **Go with Gin:** A fast and lightweight web framework for Go.
- **PostgreSQL:** A powerful, open-source object-relational database system.
- **Docker and Docker Compose:** For containerization and easy setup.
- **JWT Authentication:** Secure your routes with JSON Web Tokens.
- **GORM:** A developer-friendly ORM library for Go.
- **CRUD Operations:** A sample "Book" resource with Create, Read, Update, and Delete endpoints.
- **Best Practices:** A clean and scalable project structure.

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go](https://golang.org/doc/install) (for local development)

### Running the Application

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/go-gin-starter.git
    cd go-gin-starter
    ```

2.  **Create a `.env` file:**

    Create a `.env` file in the root of the project and add the following environment variables:

    ```env
    DB_HOST=postgres
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=postgres
    DB_NAME=go_gin_starter
    API_PORT=8080
    JWT_SECRET=your-secret-key
    ```

3.  **Run with Docker Compose:**

    ```bash
    docker-compose up --build
    ```

    The application will be running at `http://localhost:8080`.

## API Endpoints

### Authentication

-   `POST /auth/register`: Register a new user.
-   `POST /auth/login`: Login and get a JWT token.

### Books (Protected)

-   `GET /books`: Get all books.
-   `POST /books`: Create a new book.
-   `GET /books/:id`: Get a book by ID.
-   `PUT /books/:id`: Update a book by ID.
-   `DELETE /books/:id`: Delete a book by ID.

To access the protected book endpoints, you need to include the JWT token in the `Authorization` header:

```
Authorization: Bearer <your-jwt-token>
```

## Project Structure

```
.
├── cmd
│   └── api
│       └── main.go
├── config
│   └── config.go
├── internal
│   ├── controllers
│   │   ├── auth_controller.go
│   │   └── book_controller.go
│   ├── middleware
│   │   └── auth_middleware.go
│   ├── models
│   │   ├── book.go
│   │   └── user.go
│   ├── repositories
│   │   ├── book_repository.go
│   │   └── user_repository.go
│   ├── routes
│   │   ├── auth_routes.go
│   │   └── book_routes.go
│   └── services
│       ├── auth_service.go
│       └── book_service.go
├── pkg
│   ├── database
│   │   └── database.go
│   └── utils
├── .env
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

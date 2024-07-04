# Student CRUD API

This project is a simple CRUD API for managing students using Golang and the Gin framework.

## Setup Instructions

1. Clone the repository:
    ```sh
    git clone https://github.com/your-username/student-crud-api.git
    cd student-crud-api
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Run the API:
    ```sh
    make run
    ```

4. The API will be available at `http://localhost:8080`.

## Endpoints

- `POST /api/v1/students` - Add a new student
- `GET /api/v1/students` - Get all students
- `GET /api/v1/students/:id` - Get a student by ID
- `PUT /api/v1/students/:id` - Update a student
- `DELETE /api/v1/students/:id` - Delete a student
- `GET /healthcheck` - Health check endpoint

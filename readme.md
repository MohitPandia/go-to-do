# Todo Management System

A RESTful Todo Management System built using Go, GORM, and PostgreSQL to handle tasks efficiently. This project includes user management, task categorization, and the ability to set deadlines and track completion status.

## Features

1. **User Management**:
   - Create and manage users.
   - Secure user authentication and authorization.
2. **Todo Management**:
   - Create, update, delete, and retrieve todos.
   - Associate todos with specific users and categories.
   - Set deadlines and mark tasks as completed.
3. **Category Management**:
   - Define categories for tasks.
   - Associate tasks with predefined categories.
4. **Soft Deletes**:
   - Use GORM’s soft delete functionality to preserve data integrity.

---

## Technologies Used

- **Backend**: Go
- **Database**: PostgreSQL
- **ORM**: GORM
- **API Testing**: Postman

---

## API Endpoints
` By default this runs on port-8080

### Users
- `POST("/api/users/create)`
- `GET("/api/users/getAll")`
- `GET("/api/users/get-by-pid")`
- `POST("/api/users/delete")`
- `PUT("api/users/update")`

### Todos
- `POST /todos` - Create a new todo.
- `GET /todos` - Get all todos.
- `GET /todos/{id}` - Get todo by ID.
- `PUT /todos/{id}` - Update todo by ID.
- `DELETE /todos/{id}` - Soft delete a todo.

### Categories
- `POST /categories` - Create a new category.
- `GET /categories` - Get all categories.

---

## Running the Project

### Prerequisites

1. Install Go (>= 1.19).
2. Install PostgreSQL.
3. Install Postman for API testing.

### Steps to Run

1. Clone the repository:
   ```bash
   git clone (https://github.com/MohitPandia/go-to-do.git)
   cd go-to-do
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up PostgreSQL:
   - Create a database.
   - Update the database connection string in the environment variables or configuration file.

4. Run migrations:
   ```bash
   go run main.go migrate
   ```

5. Start the server:
   ```bash
   go run main.go
   ```

6. Test the API using Postman or any API client.

---

## License
This project is managed by Yash Pandia

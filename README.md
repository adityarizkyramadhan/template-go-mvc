# Golang MVC Project Template

Welcome to the Golang MVC Project Template! This repository serves as a starting point for building scalable, maintainable, and clean Go applications using the Model-View-Controller (MVC) architecture. The template is designed with flexibility in mind, integrating essential components such as Redis, PostgreSQL, Swagger, and more.

## Project Structure

The project is organized as follows:

├── controllers <br>
│ └── # Handlers for managing HTTP requests<br>
├── database<br>
│ ├── redis.go # Redis setup and connection<br>
│ └── postgre.go # PostgreSQL setup and connection<br>
├── docs<br>
│ └── # Swagger documentation files<br>
├── logs<br>
│ └── # Log files directory<br>
├── middleware<br>
│ └── # Middleware functions for request processing<br>
├── model<br>
│ └── # Model representations of database entities<br>
├── repositories<br>
│ └── # Database interaction and business logic<br>
├── routes<br>
│ └── # API route definitions<br>
├── seeder<br>
│ └── # Data seeding scripts<br>
├── storage<br>
│ └── # File storage directory<br>
├── utils<br>
│ └── # Utility functions and helpers<br>
└── main.go<br>
└── # Application entry point<br>



### 1. `controllers/`
The `controllers` directory is where you define your request handlers. These handlers are responsible for managing the flow of data between the client and the server, utilizing the services provided by the `repositories` layer.

### 2. `database/`
The `database` directory contains setup and connection logic for the databases used in the application. This template integrates both Redis and PostgreSQL to handle caching and persistent data storage, respectively.

- `redis.go`: Handles Redis setup and connection.
- `postgre.go`: Manages PostgreSQL setup and connection.

### 3. `docs/`
The `docs` directory is reserved for API documentation generated using Swagger. This allows you to maintain and share your API specifications in a standardized format, making it easier for others to understand and use your API.

### 4. `logs/`
The `logs` directory is where all log files are stored. Logging is essential for monitoring and debugging the application, and this directory keeps all logs organized and accessible.

### 5. `middleware/`
The `middleware` directory contains functions that are executed before your controllers handle a request. Middleware can be used for tasks like authentication, logging, and request validation.

### 6. `repositories/`
The `repositories` directory is where you manage interactions with the database. This layer abstracts the data access logic, allowing your application to interact with the database in a clean and organized way. Business logic that involves database operations is also placed here.

### 7. `routes/`
The `routes` directory is where you define your API routes. It acts as the mapping layer between HTTP requests and the appropriate controller functions, keeping your routing logic clean and maintainable.

### 8. `seeder/`
The `seeder` directory contains scripts for seeding the database with initial data. This is particularly useful during development and testing to ensure your application has the necessary data to run.

### 9. `storage/`
The `storage` directory is used for storing files, such as uploaded files, images, and other persistent data that doesn't belong in a database.

### 10. `utils/`
The `utils` directory houses utility functions and helpers that can be reused across the application. These are typically small, generic functions that simplify repetitive tasks.

### 11. `main.go`
The `main.go` file is the entry point of the application. It initializes the application, sets up routes, middleware, and database connections, and starts the server.

## Getting Started

To get started with this template:

1. Clone the repository:
   ```bash
   git clone https://github.com/adityarizkyramadhan/template-go-mvc
   ```
2. Install the required dependencies:
   ```bash
    go mod tidy
    ```
3. Set up your environment variables:
    - Create a `.env` file in the root directory.
    - Add your environment variables to the file. You can refer to the `.env.example` file for a list of required variables.

4. Run the application:
    ```bash
    go run main.go
    ```

## Contributing

Contributions are welcome! If you have any suggestions, bug reports, or feature requests, feel free to open an issue or create a pull request.

## License

This project is unliscensed.

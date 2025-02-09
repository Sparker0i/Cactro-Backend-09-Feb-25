# Customizable Cache API

## Overview

The **Customizable Cache API** is a RESTful service built with Go and the Gin framework. It provides endpoints to store, retrieve, and delete key-value pairs in an in-memory cache. The cache size is configurable via an environment variable, and the project is designed according to SOLID principles to ensure high maintainability, extensibility, and testability.

## Features

- **Customizable Cache Size:**  
  The maximum number of items in the cache can be configured via the `MAX_CACHE_SIZE` environment variable (default is 10).

- **RESTful Endpoints:**
  - `POST /cache`: Stores a key-value pair.
  - `GET /cache/{key}`: Retrieves the value for the specified key.
  - `DELETE /cache/{key}`: Deletes the specified key from the cache.

- **Graceful Shutdown:**  
  The standalone server implements graceful shutdown, allowing in-flight requests to complete before termination.

- **SOLID-Compliant Architecture:**  
  The project adheres to SOLID principles:
  - **SRP (Single Responsibility Principle):** Each package has a distinct responsibility.
  - **OCP (Open/Closed Principle):** The cache functionality is abstracted via a `CacheStore` interface, making it easy to extend (e.g., replacing the in-memory cache with Redis) without modifying existing code.
  - **LSP (Liskov Substitution Principle):** Any implementation of the `CacheStore` interface can be substituted without affecting the correctness of the system.
  - **ISP (Interface Segregation Principle):** Interfaces are kept minimal, exposing only what is necessary.
  - **DIP (Dependency Inversion Principle):** Higher-level modules (such as HTTP handlers) depend on abstractions rather than concrete implementations.

- **Testing:**  
  Comprehensive unit and integration tests cover the caching logic, configuration loading, HTTP endpoints, and graceful shutdown.

## Project Structure

```
customizable-cache-api/ 
|-- api/ 
| |-- index.go # Vercel serverless entry point for deployment on Vercel. 
|-- cmd/ 
│ |-- server/ 
│ |-- main.go # Standalone server entry point (includes graceful shutdown). 
|-- internal/ 
│ |-- cache/ 
│ │ |-- cache.go # Cache implementation and CacheStore interface. 
│ |-- config/ 
│ │ |-- config.go # Loads configuration from environment variables. 
│ |-- handlers/ 
│ │ |-- handlers.go # HTTP handlers for caching endpoints. 
│ |-- server/ 
│ |-- router.go # Assembles the router and injects dependencies. 
|-- go.mod # Go module definition. 
|-- vercel.json # Configuration file for Vercel deployments.
```


## Design and SOLID Principles

- **Cache Package:**  
  - **Responsibility:** Provides caching logic.
  - **Abstraction:** Defines a `CacheStore` interface.
  - **Extensibility:** In-memory implementation can be replaced (e.g., with Redis) without affecting other modules.
  
- **Configuration Package:**  
  - **Responsibility:** Loads settings (like `MAX_CACHE_SIZE`) from environment variables.
  
- **HTTP Handlers:**  
  - **Responsibility:** Handle API requests.
  - **Dependency:** Depend on the `CacheStore` abstraction (promoting dependency inversion).
  
- **Router/Server Package:**  
  - **Responsibility:** Assembles the application by wiring dependencies together and registering endpoints.


- **Vercel Deployment:**  
  - **Responsibility:** Provides a serverless function entry point for deployment on Vercel.


## Running the Application

### Local Standalone Server

1. **Install Go:**  
   Ensure Go is installed. [Download Go](https://golang.org/dl/)

2. **Set the Environment Variable (Optional):**  
   To customize the cache size, set the `MAX_CACHE_SIZE` environment variable:
   ```bash
   export MAX_CACHE_SIZE=20
   ```

   If not set, the default cache size is 10.

3. Run the Server:

   ```bash
   go run cmd/server/main.go
   ```

   The API will be available at: http://localhost:8080

## Testing

To run all tests, execute:

```bash
go test ./...
```

This command runs unit tests for the cache, configuration, HTTP handlers.

## Conclusion

This SOLID-compliant, customizable caching API is designed for robustness, maintainability, and ease of extension. It features clear separation of concerns, dependency injection, and a comprehensive test suite. Whether deployed as a standalone server or as a serverless function on Vercel, this API serves as a strong foundation for further development.

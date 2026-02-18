# Wallet Microservice

## Tech Stack

- **Language**: Go
- **Framework**: Gin
- **Database**: PostgreSQL
- **Containerization**: Docker & Docker Compose

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Postman (Optional, for API testing)

### Running with Docker Compose

The simplest way to run the application and its database is using Docker Compose.

1.  Ensure Docker is running on your machine.
2.  From the project root, run the following command:
    ```bash
    docker compose up --build
    ```
    This will build the Go application, start the PostgreSQL database, run the initial schema migrations from `init.sql`, and start the API server on `http://localhost:8080`.

### Importing the Postman Collection

1.  Open Postman.
2.  Click on `File` > `Import...`.
3.  Select the `wallet_microservice.postman_collection.json` file from the root of this project.
4.  The "Wallet Microservice" collection will be imported. It contains pre-configured requests for `Get Balance` and `Withdraw Funds` that are ready to use with the running Docker container.

## Decisions in Project Building

### Locking Mechanism

I chose an **optimistic locking** mechanism for database updates.
From the requirement "A user has a balance in a digital wallet application," I inferred that the system is user-centric. Since a single wallet is unlikely to have high-frequency concurrent writes from multiple sources, optimistic locking is efficient.

To handle the rare cases of collision (concurrent updates), the service layer implements a **retry mechanism** (defaulting to 3 attempts). If the version check fails, the operation is retried automatically before returning an error to the user.

### GRPC vs REST API

I chose to use **REST API** instead of gRPC.
The requirements state: "A user has a balance... and wants to perform transactions." This implies direct interaction from a frontend or client application. REST is universally supported and easier to integrate for web/mobile clients compared to gRPC, which is often better suited for internal service-to-service communication.

Given the current monolithic scope, REST minimizes complexity.

### Error Handling

Domain errors are mapped to specific HTTP status codes to provide meaningful feedback to the client:

- **409 Conflict**: Returned when a concurrent update occurs and all retries fail.
- **422 Unprocessable Entity**: Returned for business logic errors, such as insufficient funds.
- **404 Not Found**: Returned when the wallet or user does not exist.

### Database Schema Decisions

#### No Password in User Table

Since this project focuses on the wallet service logic, I omitted authentication details (like passwords) from the `users` table. Implementing a full auth system is outside the scope of this specific task and would slow down the development of the core wallet features.

#### `updated_at` Columns

To keep the schema lean, `updated_at` is only present in the `wallets` table, as that is the primary mutable entity. `users` and `transactions` are treated as immutable or append-only for this iteration.

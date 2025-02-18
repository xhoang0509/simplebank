# Simple Bank

This project is a simple banking system implemented in Go. It includes functionalities for managing accounts and transfers, with tests to ensure the correctness of the operations.

## Features

- Create, update, and retrieve accounts
- Create and retrieve transfers
- List accounts and transfers
- Concurrent transfer transactions with deadlock handling

## Technologies

- Go
- SQL
- `sqlc` for generating type-safe Go code from SQL queries
- `testify` for testing

## Setup

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/simplebank.git
   cd simplebank
   ```
2. **Install dependencies:**
   ```sh
   go mod tidy
   ```
3. **Set up the database:**
   - Create a PostgreSQL database.
   - Update the database connection settings in the `db/sqlc.yaml` file.
4. **Generate Go code from SQL:**
   ```sh
   sqlc generate
   ```
5. **Run the tests:**
   ```sh
   make test
   ```
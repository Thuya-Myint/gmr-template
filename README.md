
# gmr template

-A production-oriented Go backend template with MongoDB and Redis, featuring a clean architecture, reliable environment configuration, and solutions for common MongoDB driver DNS issues across local and cloud deployments
## Installation

### Prerequisites
- Go 1.21+
- MongoDB
- Redis

### Steps

1. Install dependencies
    ```bash
    go mod tidy

2.	Configure environment variables.  

      Create a .env file and define the required variables.

      The list of required environment variables can be found in:
    ```bash
    internal/config/config.go

3. Build the application
    ```bash
    go build ./cmd/api

4. Run the application
    ```bash
    go run ./cmd/api
## After running

Once the application is running, it will automatically attempt to connect
to MongoDB and Redis using the configured environment variables.

Successful connections will be logged to the console.

If startup fails, ensure MongoDB and Redis are running and that all required
environment variables are correctly configured.
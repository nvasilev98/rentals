# Rentals

### Run Unit Tests
```bash
make unit-tests
```

### Run application locally
1. Export Env Variables

    ```bash
    export HOST=<host>

    export PORT=<port>

    export DB_HOST=<db-host>

    export DB_PORT=<db-port>

    export DB_USERNAME=<db-user>

    export DB_PASSWORD=<db-pass>

    export DB_NAME=<db-name>
    ```
2. Spin up database in a separate session

    ```bash
    docker compose up
    ```
    
3. Run application

    ```bash
    go run cmd/rentals/main.go
    ```

# REST API GOLANG
## Setup and Run 

### Prerequisites

- Go 1.20 or higher
- MySQL server

### Database Setup
1. Create a MySQL database named `db_golang`:
   ```sql
   CREATE DATABASE db_golang;
   ```
2. import sql on file database.sql

3. edit env.example file to .env with database_name = db_golang 

### Running the Application

1. Install dependencies:
   ```bash
   go mod init golang-api
   go mod tidy
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

3. Access the API at `http://localhost:9800`

### Running Tests
before run test please edit on file user and message test username:password database internal
```bash
go test tests\user_test.go
go test tests\message_test.go
```


1. Install Prerequisites
Install Go
* Download and install Go from the official website.
* Verify the installation:


Install MySQL
- Download and install MySQL from the official website.
- Start the MySQL server and create a database for the project:
```sql
CREATE DATABASE irctc;
```

2. Clone the Project

```bash
    git clone https://github.com/rimo02/railway-management-system.git
```
3. Configure the Environment
Update Database Credentials
In the `database/db.go` file, update the dsn (Data Source Name) with your database credentials:

```go
    dsn := "user:password@tcp(127.0.0.1:3306)/irctc?charset=utf8mb4&parseTime=True&loc=Local"
```
Replace user with your MySQL username, and password with your MySQL password.

---
4. Install Dependencies
Install the required Go packages:


```bash
    go mod tidy
```

5. Run the Application
Start the Go server:

```bash
    go run main.go
```
---
1. Test the API
Use tools like Postman or curl to test the endpoints. Below are some examples.

#### Register a User
```bash
    curl -X POST http://localhost:8080/api/register \
    -H "Content-Type: application/json" \
    -d '{"username": "user1", "password": "pass123", "role": "user"}'
```
#### Login
```bash
    curl -X POST http://localhost:8080/api/login \
    -H "Content-Type: application/json" \
    -d '{"username": "user1", "password": "pass123"}'
```
Response will include a token:

```json
    {
        "token": "your_jwt_token"
    }
```
#### Add a Train (Admin)
```bash
    curl -X POST http://localhost:8080/api/admin/train \
    -H "Content-Type: application/json" \
    -H "X-API-Key: your_admin_api_key" \
    -d '{"source": "A", "destination": "B", "total_seats": 100, "available_seats": 100}'
```

#### Book a Seat
```bash
    curl -X POST http://localhost:8080/api/book \
    -H "Content-Type: application/json" \
    -H "Authorization: your_jwt_token" \
    -d '{"train_id": 1}'
```

6. Verify Database Changes
Use a MySQL client to check the users, trains, and bookings tables for updates:

```sql
    SELECT * FROM users;
    SELECT * FROM trains;
    SELECT * FROM bookings;
```
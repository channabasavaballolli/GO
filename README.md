



Follow these steps to set up the PetClinic API on your system.


 Create PostgreSQL Database & Tables**

Run these SQL commands inside `psql`:

```sql
CREATE DATABASE petclinic;

\c petclinic;

CREATE TABLE owners (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    contact VARCHAR(20),
    email VARCHAR(100)
);

CREATE TABLE pets (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    species VARCHAR(50),
    age INT,
    owner_id INT REFERENCES owners(id) ON DELETE CASCADE
);

CREATE TABLE appointments (
    id SERIAL PRIMARY KEY,
    pet_id INT REFERENCES pets(id) ON DELETE CASCADE,
    appointment_date DATE,
    description TEXT
);
```

---

2) Add Environment Variables (.env)**

Create a `.env` file in your project root:

```
PG_USER=postgres
PG_PASSWORD=1234
PG_HOST=localhost
PG_PORT=5432
PG_DB=petclinic
PG_SSLMODE=disable

JWT_SECRET=MySuperSecretKey123
JWT_EXPIRE_HOURS=2
```

---

3) Install Go Modules and Dependencies

go mod init petclinic
go get github.com/gorilla/mux
go get github.com/lib/pq
go get github.com/golang-jwt/jwt/v5
go get github.com/sirupsen/logrus
go get github.com/joho/godotenv
```

---

Run the Server


go run .
```

If successful, you will see:

```
Connected to PostgreSQL successfully!
Server running at http://localhost:8080
```

---

*Authentication (JWT)
Login (Public)

Get a JWT token:

```
POST http://localhost:8080/login
```

Body:

```json
{
  "username": "Beast",
  "password": "Channu@4321"
}
```

*Response:

```json
{
  "token": "<JWT_TOKEN_HERE>"
}
```

Copy the token.

---

Using JWT for Protected Routes

Add this header to ALL `/api/...` requests:

```
Authorization: Bearer <token>
```

Example:

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

---
API Endpoints

All protected endpoints start with `/api`.

---

Owners


Create Owner

POST`/api/owners`

```json
{
  "name":"Alice",
  "contact":"999999",
  "email":"alice@example.com"
}
```

Get All Owners

GET `/api/owners`

 Get Owner by ID

GET `/api/owners/{id}`

 Update Owner

PUT `/api/owners/{id}`

```json
{
  "name":"Updated Name",
  "contact":"888888",
  "email":"updated@example.com"
}
```

Delete Owner

DELETE `/api/owners/{id}`

---

#Pets

#Create Pet

POST `/api/pets`

```json
{
  "name":"Rex",
  "species":"Dog",
  "age":3,
  "owner_id":1
}
```

Get All Pets

GET `/api/pets`

 Get Pet by ID

GET `/api/pets/{id}`

Update Pet

PUT `/api/pets/{id}`

```json
{
  "name":"Max",
  "species":"Dog",
  "age":5,
  "owner_id":1
}
```

Delete Pet

DELETE `/api/pets/{id}`

---

 Appointments

Create Appointment

POST `/api/appointments`

```json
{
  "pet_id":1,
  "appointment_date":"2025-11-10",
  "description":"Vaccination"
}
```

 Get All Appointments

GET `/api/appointments`

 Get Appointment by ID

GET `/api/appointments/{id}`

Update Appointment

PUT `/api/appointments/{id}`

```json
{
  "pet_id":1,
  "appointment_date":"2025-11-20",
  "description":"Updated Description"
}
```

Delete Appointment

DELETE `/api/appointments/{id}`

---





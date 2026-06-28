# 🚗 SpotSync API

Smart Parking & EV Charging Reservation System built with Go, Echo, GORM, PostgreSQL, and JWT authentication.

## 🌍 Live Demo

* **Backend API:** https://your-deployment-url.com
---

## 📌 Project Overview

SpotSync is a centralized parking management platform designed for busy airports and malls. The system supports regular parking and EV charging reservations while ensuring that parking zones never exceed their capacity, even under concurrent booking requests.

The project implements Clean Architecture principles and uses database transactions with row-level locking (`FOR UPDATE`) to prevent race conditions during reservations.

---

## ✨ Features

### 🔐 Authentication

* User registration
* User login
* JWT authentication
* Password hashing with bcrypt
* Role-based authorization

### 🚗 Parking Zones

* Create parking zones (Admin)
* View all parking zones
* View single parking zone
* Dynamic available spot calculation

### 📅 Reservations

* Reserve parking spots
* View personal reservations
* Cancel reservations
* View all reservations (Admin)
* Concurrency-safe booking system using transactions and row locking

---

## 🛠 Tech Stack

| Technology | Usage              |
| ---------- | ------------------ |
| Go 1.24+   | Backend language   |
| Echo       | HTTP framework     |
| GORM       | ORM                |
| PostgreSQL | Database           |
| JWT        | Authentication     |
| bcrypt     | Password hashing   |
| Validator  | Request validation |

---

## 🏗 Domain Driven Design

The project follows Domain Driven Design principles:

```text
SpotSync-server/
│
├── cmd/
│   └── main.go
│
├── auth/
│   └── jwt.go
│
├── config/
│   └── config.go
│
├── database/
│   └── postgres.go
│
├── domain/
│   │
│   ├── users/
│   │   ├── dto/
│   │   │   ├── request.go
│   │   │   └── response.go
│   │   ├── entity.go          # User model
│   │   ├── repository.go
│   │   ├── service.go
│   │   ├── handler.go
│   │   └── routes.go
│   │
│   ├── parking_zones/
│   │   ├── dto/
│   │   │   ├── request.go
│   │   │   └── response.go
│   │   ├── entity.go          # ParkingZone model
│   │   ├── repository.go
│   │   ├── service.go
│   │   ├── handler.go
│   │   └── routes.go
│   │
│   └── reservations/
│       ├── dto/
│       │   ├── request.go
│       │   └── response.go
│       ├── entity.go          # Reservation model
│       ├── repository.go
│       ├── service.go
│       ├── handler.go
│       └── routes.go
│
├── middleware/
│   ├── auth.go
│   └── admin.go
│
├── server/
│   └── http.go
│
├── utils/
│   ├── error.go
│   └── parse.go
│
├── .env
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

### Layer Responsibilities

| Layer      | Responsibility                  |
| ---------- | ------------------------------- |
| DTO        | Request and response structures |
| Handler    | HTTP request handling           |
| Service    | Business logic                  |
| Repository | Database operations             |
| Models     | Database entities               |

Dependencies flow in one direction:

```text
Handler
   ↓
Service
   ↓
Repository
   ↓
Database
```

---

## 🔒 Concurrency Handling

The reservation system prevents overbooking by using:

* Database Transactions
* Row-Level Locking (`FOR UPDATE`)
* Atomic reservation creation

This guarantees that two users cannot reserve the last available spot simultaneously.

Example:

```go
tx.Clauses(
    clause.Locking{
        Strength: "UPDATE",
    },
).First(&zone, zoneID)
```

---

## ⚙️ Environment Variables

Create a `.env` file:

```env
DSN=dataBase_connection
PORT=8080
JWT_SECRET=your-super-secret-key
```
---

## 🚀 Running Locally

### 1. Clone the Repository

```bash
git clone https://github.com/noboKumar/SpotSync-server.git

cd SpotSync-server
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Create Environment Variables

```bash
cp .env
```

Update the values inside `.env`.

### 5. Run the Application

```bash
go run cmd/main.go
```

or

```bash
go run main.go
```

depending on your project structure.

The server will start at:

```text
http://localhost:8080
```

---

## 🔥 Development Mode (Hot Reload)

Install Air:

```bash
go install github.com/air-verse/air@latest
```

Run:

```bash
air
```

---

## 📚 API Endpoints

### Authentication

| Method | Endpoint                | Access |
| ------ | ----------------------- | ------ |
| POST   | `/api/v1/auth/register` | Public |
| POST   | `/api/v1/auth/login`    | Public |

---

### Parking Zones

| Method | Endpoint            | Access |
| ------ | ------------------- | ------ |
| POST   | `/api/v1/zones`     | Admin  |
| GET    | `/api/v1/zones`     | Public |
| GET    | `/api/v1/zones/:id` | Public |

---

### Reservations

| Method | Endpoint                               | Access        |
| ------ | -------------------------------------- | ------------- |
| POST   | `/api/v1/reservations`                 | Authenticated |
| GET    | `/api/v1/reservations/my-reservations` | Authenticated |
| DELETE | `/api/v1/reservations/:id`             | Authenticated |
| GET    | `/api/v1/reservations`                 | Admin         |

---

## 🧪 Testing API

You can test the API using:

* Postman
* Thunder Client
* Insomnia
* cURL

Example:

```bash
curl http://localhost:8080/api/v1/zones
```

---

## 📜 License

This project was developed as part of an academic assignment and is intended for educational purposes.

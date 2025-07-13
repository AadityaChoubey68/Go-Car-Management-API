# 🚗 CarZone - Car Management REST API

CarZone is a backend REST API written in Go (Golang) that allows you to manage a collection of cars and their associated engine specifications. This project supports CRUD operations, integrates with PostgreSQL for data persistence, and is containerized using Docker.

## 📦 Features

- 🚀 Fast RESTful API using Go
- 🗃️ PostgreSQL as the primary database
- 🔁 Full CRUD for `cars` and their linked `engines`
- 🐳 Dockerized for easy local setup
- 📄 Schema auto-execution on startup
- 🔒 Validation for fuel types
- 🔄 Transaction-safe database operations

---

## 🛠️ Tech Stack

- **Golang** (net/http, database/sql)
- **PostgreSQL** (via `lib/pq`)
- **Docker + Docker Compose**
- **UUID** for object identification
- **Chi Router** for clean routing

---

## 📁 Project Structure

carzone/
├── Dockerfile
├── docker-compose.yml
├── main.go
├── db/
│ └── schema.sql
├── models/
│ └── car.go
├── store/
│ └── car_store.go
├── handlers/
│ └── car_handler.go
└── utils/
└── validate.go


---

## 🚀 Getting Started

### Prerequisites

- [Go](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/)

---

### 🐳 Run using Docker (Recommended)

```bash
docker-compose up --build

🧪 Sample Car Request Body
{
  "id": "a3f86b1a-1111-4c19-bae2-fc60329d01234",
  "name": "Audi A4",
  "year": "2023",
  "brand": "Audi",
  "fuel_type": "petrol",
  "engine": {
    "eng_id": "9746be12-07b7-42a3-b8ab-7d1f209b63d7",
    "displacement": 3000,
    "no_of_cylinders": 1,
    "car_range": 2000
  },
  "price": 38000.00
}


🛣️ API Endpoints
Method	Endpoint	  Description
GET/cars	        List all cars
GET/cars/{id}	    Get car by ID
POST/cars	        Create a new car
PUT/cars/{id}	    Update an existing car
DELETE/cars/{id}	Delete a car

🧹 To Reset Database

docker-compose down -v
docker-compose up --build

🧑‍💻 Author
Aditya N. Chaubey
B.Tech (GGSIPU, 2021–2025)
Skills: React.js, Next.js, Golang, REST APIs, PostgreSQL

📄 License
This project is open-source and free to use under the MIT License.

---

Let me know if you’d like to:
- Add Swagger/OpenAPI documentation
- Include Postman collection
- Or turn this into a full-stack app with frontend too!


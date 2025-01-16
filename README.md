```markdown
# MemeCoin API

MemeCoin API is a RESTful service for managing meme coins. The application provides endpoints to create, fetch, update, delete, and "poke" (increment popularity) meme coins. It uses MongoDB as the database and is built with Go using the Gin framework.

---

## Features

- Add new meme coins with unique names and descriptions.
- Retrieve details of a meme coin by its ID.
- Update the description of a meme coin.
- Delete a meme coin from the database.
- Poke a meme coin to increment its popularity score.

---

## Project Structure

MemeCoin_api/
├── api/
│   ├── controllers/                    # Handlers for API endpoints
│   │   ├── meme_coin_controller.go
│   ├── models/                         # Data models and MongoDB schema
│   │   ├── meme_coin_model.go
│   ├── routes/                         # API route definitions
│   │   ├── routes.go
│   ├── services/                       # Services package
│   │   ├── meme_coin_service.go
├── config/
│   ├── db.go                           # MongoDB connection setup
├── Dockerfile                          # Dockerfile for building the application image
├── docker-compose.yml                  # Docker Compose configuration
├── go.mod                              # Go module file
├── go.sum                              # Dependencies checksum
├── main.go                             # Application entry point
├── README.md                           # Project documentation
```

---

## Prerequisites

- **Go**: Version 1.20+
- **Docker**: Version 20.10+
- **Docker Compose**: Version 2.31+
- **MongoDB**: Version 5.0+ (optional for local setup)

---

## Running the Application

### **1. Local Setup**

#### **Clone the repository**
```bash
git clone https://github.com/cheildo/meme_coin_api.git
cd MemeCoin_api
```

#### **Install Dependencies**
```bash
go mod tidy
```

#### **Run MongoDB Locally**
Make sure MongoDB is running on `localhost:27017`.

#### **Start the Application**
```bash
go run main.go
```

#### **Access the API**
The server runs on `http://localhost:8080`.

---

### **2. Running with Docker**

#### **Build and Start the Containers**
```bash
docker compose up --build
```

#### **Access the API**
- The application will be available at `http://localhost:8080`.
- MongoDB will be accessible at `localhost:27017`.

---

## API Endpoints

### **Base URL**: `http://localhost:8080/meme_coins/api/`

| Method | Endpoint               | Description                     |
|--------|------------------------|---------------------------------|
| POST   | `/`                    | Create a new meme coin.         |
| GET    | `/:id`                 | Fetch a meme coin by ID.        |
| PUT    | `/:id`                 | Update a meme coin description. |
| DELETE | `/:id`                 | Delete a meme coin by ID.       |
| POST   | `/:id/poke`            | Increment a meme coin's score.  |

---

## Configuration

### **Environment Variables**
| Variable      | Description                           |   Docker value             | Local value                |
|---------------|---------------------------------------|----------------------------|----------------------------|
| `MONGO_URI`   | Docker MongoDB connection string.     | `mongodb://mongo:27017`    | `mongodb://localhost:27017`|

You can define this variable in a `.env` file or directly in `docker-compose.yml`.

---

## Notes

- Ensure MongoDB is running and accessible when starting the application locally.
- Docker Compose automatically creates a virtual network for containers to communicate with each other.


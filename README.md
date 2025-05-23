# 🛒 Karting Challenge (Go)

This is a robust Go-based backend implementation of the **Karting Challenge** API, based on the [OpenAPI 3.1 specification](https://swagger.io/specification/).

> ✅ Built with production-grade practices for clean architecture, performance, and testability.

---

## 📦 Features

- Fully compliant with the [OpenAPI spec](https://orderfoodonline.deno.dev/public/openapi.yaml)
- Serve endpoints to:
  - List products
  - Get product by ID
  - Place an order (with promo code support)
- Promo code validation from 3 large `.gz` files
- Promo codes must appear in **at least 2 out of 3** files to be valid
- Files are:
  - Downloaded automatically at startup
  - Parsed and cached in-memory for high-speed lookups
  - Periodically refreshed in the background using goroutines -- extension, in progress
- API key-based authentication
- Environment-based configuration via `godotenv`

---

## 📁 Project Structure
```
├── cmd/
│ └── main.go # App entry point
├── internal/
│ ├── data/ # Storage for large promo files (if downloaded)
│ ├── handlers/ # Route handlers (product, order)
│ ├── middleware/ # Custom middleware (e.g. API key auth)
│ ├── models/ # Data models (Product, Order, etc.)
│ ├── router/ # Gin router setup
│ ├── service/ # Business logic and promo validator
│ └── utils/ # Helpers (e.g. file downloader)
├── client #Frontend code, react app
├── go.mod
├── .env # Environment variables
└── README.md
```

## 🚀 Getting Started

### 1. Clone the repo
```
git clone https://github.com/mani87/karting.git
cd karting-challenge
```

### 2. Start the app
```
cd cmd/server
go run main.go
your server is running at port 8080
```

### 3. Start testing it via Postman
Here is a sample curl that can be used;
```
curl --location 'localhost:8080/api/order' \
--header 'x-api-key: create_order' \
--header 'Content-Type: application/json' \
--data '{
    "couponCode": "HAPPYHR",
    "items": [
        {
            "productId": "10",
            "quantity": 3
        }
    ]
}'
```
It looks like this for now, Frontend is in-progress for now;
<img width="1450" alt="Screenshot 2025-05-24 at 12 29 29 AM" src="https://github.com/user-attachments/assets/06e3eba2-a436-498b-9102-42342f76d3b3" />




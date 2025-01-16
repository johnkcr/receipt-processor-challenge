# Receipt Processor

A simple API to process receipts and calculate points based on specific rules. This application is written in Go and uses the `chi` router.

---

## Prerequisites

- [Go 1.20+](https://go.dev/dl/)
- [Docker](https://www.docker.com/)
- [cURL](https://curl.se/) or a REST client like [Postman](https://www.postman.com/).

---

## Getting Started

git clone https://github.com/your-username/receipt-processor.git
cd receipt-processor

---

## Running the Application

### Option 1: Using Go

go mod tidy
go run main.go

The application will run on http://localhost:3000

---

### Option 2: Using Docker

docker build -t receipt-processor .
docker run -p 3000:3000 receipt-processor

The application will run on http://localhost:3000

---

### Option 3: Using Docker Compose (Optional)

docker-compose up --build

The application will run on http://localhost:3000

---

## Endpoints

### Submit a Receipt

curl -X POST http://localhost:3000/receipts/process \
-H "Content-Type: application/json" \
-d '{
      "retailer": "Target",
      "purchaseDate": "2022-01-01",
      "purchaseTime": "13:01",
      "items": [
          {"shortDescription": "Mountain Dew 12PK", "price": "6.49"}
      ],
      "total": "6.49"
    }'

Response:

{
  "id": "unique-receipt-id"
}

---

### Get Receipt Points

curl http://localhost:3000/receipts/unique-receipt-id/points

Response:

{
  "points": 50
}

---

## Development

### Code Generation

oapi-codegen -generate "chi-server" -package gen -o api/gen/handlers_gen.go api/spec/api.yaml
oapi-codegen -generate "types" -package gen -o api/gen/types.go api/spec/api.yaml

---

## Troubleshooting

### GLIBC Version Issues in Docker

If you encounter errors like:

./receipt-processor: /lib/aarch64-linux-gnu/libc.so.6: version `GLIBC_2.34' not found (required by ./receipt-processor)

Solutions:

1. Use a Compatible Base Image: Switch to a base image like `debian:bookworm-slim` which includes `glibc 2.36`.

2. Manually Install glibc: Update the Dockerfile to download and install a compatible version of glibc.

3. Build a Static Binary: Use CGO_ENABLED=0 to create a binary that doesn't depend on glibc.

---

## License

This project is licensed under the MIT License. See the LICENSE file for details.
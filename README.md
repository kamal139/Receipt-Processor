# Receipt-Processor

## Overview

The Receipt Processor is a Go web service that processes receipts and calculates points based on specific rules.

## Requirements

- Go 1.17 or later
- Docker (optional, for containerized deployment)

## Cloning and Setup

1. **Clone the Repository:**

    ```bash
    git clone https://github.com/kamal39/receipt-processor.git
    cd receipt-processor
    ```

2. **Build and Run with Docker:**

    ```bash
    docker build -t receipt-processor .
    docker run -p 8080:8080 receipt-processor
    ```

    Note: Docker is optional and can be used for containerized deployment.

## API Endpoints

### Process Receipts (POST)

- **Endpoint:** `/receipts/process`

    ```bash
    # Invoke Style
    generated_id=$(curl -X POST -H "Content-Type: application/json" -d '@example/sample_receipt.json' http://localhost:8080/receipts/process | jq -r '.id')

    # Revoke Style
    curl -X DELETE http://localhost:8080/receipts/$generated_id
    ```

    **Curl Style:**
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"retailer": "ExampleMart", "total": 50.25, "items": ["item1", "item2"], "purchaseDate": "2024-01-08T15:30:00Z"}' http://localhost:8080/receipts/process
    ```

### Get Points (GET)

- **Endpoint:** `/receipts/{id}/points`

    ```bash
    # Invoke Style
    curl http://localhost:8080/receipts/{id}/points

    # Revoke Style
    curl -X DELETE http://localhost:8080/receipts/{id}/points
    ```

    **Curl Style:**
    ```bash
    # Replace {id} with the actual ID obtained from the /receipts/process endpoint
    curl http://localhost:8080/receipts/{id}/points
    ```

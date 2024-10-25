# urlshortener

## Overview

`urlshortener` is a simple URL shortening service implemented in Go. It provides endpoints to shorten URLs and retrieve the original URLs using the shortened versions. The service also includes Prometheus metrics for monitoring.

## Prerequisites

- Go 1.23+
- Docker & docker-compose
- `make` for running commands

## Setup

1. **Clone the repository**:
   ```sh
   git clone https://github.com/yourusername/urlshortener.git
   cd urlshortener
   ```

2. **Install dependencies**:
   ```sh
   go mod tidy
    ```

## Run the Application
To start the application, run:
   ```sh
   make up
   ```

The database will start listening on port 5432.
The server will start on port 8000.

## Endpoints
    POST /{url}: Shorten a URL.
    GET /{shortened_url}: Retrieve the original URL using the shortened version.
    GET /metrics: Prometheus metrics endpoint.

## Testing

Several HTTP Requests are already set in tests.http file ; the POST request store the result (the tiny url path) as parameter, enabling the call to the GET request.


## Metrics
PromHTTP default values are available on /metrics ; a counter realted to request parameters could be added here.
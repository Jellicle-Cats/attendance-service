# README

## How to Start

### Prerequisites

- Visual Studio Code (VSCode) or any text editor.
- Docker for running the server component.
- Go programming language for the consumer and producer components. Ensure Go is installed.

### Getting Started

1. Open your code editor or workspace.

2. Navigate to the `/server` directory in your terminal or command prompt:

```sh
cd server
```

3. Run the server using Docker Compose:

```sh
docker compose up
```

4. Open another terminal or command prompt and navigate to the `/consumer` directory:

```sh
cd consumer
```

5. Set up and run the consumer component:

```sh
go get
go mod tidy
go run .
```

6. Open yet another terminal or command prompt and navigate to the `/producer` directory:

```sh
cd producer
```

7. Set up and run the producer component:

```sh
go get
go mod tidy
go run .
```

8. The REST API runs on port 3456.

## API Endpoints

### POST /checkin or /checkout

Use these endpoints to check in or check out a user from a booking. Make a POST request with the following JSON data:

```json
{
  "UserID": 3,
  "BookingID": 4
}
```

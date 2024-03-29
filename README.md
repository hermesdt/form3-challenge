## Summary

This projects is a simple CRUD implementation of a server responsible for receiving payments from different banks.

### Endpoints

```
GET /v1/payments - List payments
GET /v1/payments/some-id - Show a specific payment
POST /v1/payments - Create a new payment
PUT /v1/payments/some-id - Updates a specific payment
DELETE /v1/payments/some-id - Deletes a specific payment
```

## Requisites

As the project uses mongodb as database you can start it with:

`docker-compose up`

## Start the application

You might want to copy and modify `.env.test` file to `.env` at your convenience as the projects needs those env vars to start. Then you can load the variables with `source .env`.

Then start the server with:

`go run cmd/server.go`

## Run the tests

To run the tests you must have mongodb running and correctly configured the environment variables for the tests:

1) Start mongodb: `docker-compose up`
2) Load env vars: `source .env.test`
3) Run tests: `go test ./...`

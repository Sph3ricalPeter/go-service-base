# Simple GO Service

A simple (almost minimal) service written in GO with a multistage docker build &amp; 15MB container size

## Usage

Simply run the following command to start the service:

```shell
docker compose up
```

There's a single endpoint `/alive`:

```shell
curl http://localhost:8080/alive
# prints {"message":"Yo I'm alive!"}
```

Slog is used for logging. Calling the endpoint will print:

```shell
go-first-try  | [GIN] 2024/03/06 - 18:25:50 | 200 |     104.666µs |    192.168.65.1 | GET      "/alive"
```
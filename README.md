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
# prints out {"message":"Yo I'm alive!"}
```

Calling the endpoint will print:

```text
go-base-app  | [GIN] 2024/03/06 - 18:37:39 | 200 |     142.333µs |    192.168.65.1 | GET      "/alive"
```
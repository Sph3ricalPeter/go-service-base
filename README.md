# Simple GO Service

A simple (almost minimal) service written in GO with a multistage docker build &amp; 15MB container size

## Usage

Simply run the following command to start the service:

```shell
docker compose up --build
```

There's an endpoint `/alive`:

```shell
curl http://localhost:8080/alive
# prints out '{"message":"Yo I'm alive!"}'
```

Calling the endpoint will print:

```text
2024/10/10 22:55:27 "GET http://localhost:8080/alive HTTP/1.1" from 172.18.0.1:60952 - 200 24B in 40.452µs
```

Endpoint `/private` requires an api key header, default is `123123`

```shell
curl -X GET "localhost:8080/private" -w "%{http_code}"
# prints out '401'

curl -X GET "localhost:8080/private" -H "X-API-KEY: 123123"
# prints out '{"message":"I'm private!"}'
```
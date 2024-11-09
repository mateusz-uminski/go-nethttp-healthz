# go-nethttp-healthz

[![verify](https://github.com/mateusz-uminski/go-nethttp-healthz/actions/workflows/verify.yml/badge.svg)](https://github.com/mateusz-uminski/go-nethttp-healthz/actions/workflows/verify.yml)

This is a lightweight Go web application built using the standard `net/http` that provides a simple health check API endpoint. The application exposes a `/api/v1/healthz` endpoint that returns the service's health status. The endpoint can be used to verify if the application is running and operational.

Besides the README.md further documentation can be found in commits, code comments and nested README files.

Feel free to explore and copy everything you want. Enjoy!


# Usage

## Build the application

```sh
make build
```

## Run the application

```sh
make build
make run
```

## Run tests

```sh
make tests
```

## Build a docker image

```sh
docker build -t go-nethttp-healthz .
```

## Run the docker image

```sh
docker run -it --rm -p 8080:8080 -e APP_HOST=0.0.0.0 go-nethttp-healthz
```

## Example query

```sh
curl -v -XGET 127.0.0.1:8080/api/v1/healthz
```

FROM golang:1.20-bullseye AS build

WORKDIR /app
COPY . .
RUN make build

FROM scratch

WORKDIR /app
COPY --from=build /app/build/main .
ENTRYPOINT ["./main"]

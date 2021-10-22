FROM golang:1.17.1-buster as build

WORKDIR /app

COPY . ./
RUN go mod download
RUN go build -o server cmd/server/main.go

FROM debian:buster-slim

WORKDIR /

COPY --from=build /app/server /server

EXPOSE 8080

ENTRYPOINT ["/server"]
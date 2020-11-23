## Setup Development Server

Install NATS using Docker

```
$ docker pull nats
$ docker run -d --name nats-main -p 4222:4222 -p 6222:6222 -p 8222:8222 nats
```

Clone this repo and install dependencies.

```
$ go get -u ./...
```

To run locally

```
$ go run main.go
```

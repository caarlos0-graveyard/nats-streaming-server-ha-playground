# nats-streaming-server HA playground

## cluster

just run:

```shell
(cd cluster && docker-compose up)
```

## fault tolerant

just run:

```shell
(cd ft && docker-compose up)
```

## consumer/producer

### Producer:

```sh
go run producer/main.go CLIENT_ID nats://localhost:4221 nats://localhost:4222 nats://localhost:4223
```

### Consumer

```sh
go run consumer/main.go CLIENT_ID nats://localhost:4221 nats://localhost:4222 nats://localhost:4223
```

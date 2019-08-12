# echo-gorm-example
Example of Echo + GORM

## Usage
```sh
$ cd db
$ docker-compose up -d
$ cd ../
$ go run server.go
```

Open [localhost:1323](http://localhost:1323)

## Test
```sh
$ go test ./... -cover
```

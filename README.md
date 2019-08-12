# echo-gorm-example
Example of Echo + GORM

## Usage
```sh
$ cd db
$ docker-compose up -d
$ cd ../
$ go run server.go
```

API is available on [localhost:1323](http://localhost:1323)

### Example
#### Create
```sh
$ curl -X POST -H 'Content-Type: application/json' -d '{"name":"John Smith","language_id":1}' localhost:1323/users
```

#### Read
```sh
$ curl localhost:1323/users/1
```

#### Update
```sh
$ curl -X PUT -H 'Content-Type: application/json' -d '{"name":"John Smith","language_id":1}' localhost:1323/users/1
```

#### Delete
```Sh
$ curl -X DELETE localhost:1323/users/1
```

## Test
```sh
$ go test ./... -cover
```

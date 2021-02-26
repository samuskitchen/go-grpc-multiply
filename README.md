# gRPC Multiply
This is a project for a challenge in Golang. This api is in charge of multiplying two values, and it is driving gRPC

The following technologies were used in this project:
- [Golang 1.15](https://golang.org/dl/)
- [protobuf](https://github.com/golang/protobuf)
- [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
- [google-genproto](https://pkg.go.dev/google.golang.org/genproto)
- [google-grpc](https://pkg.go.dev/google.golang.org/grpc)
- [google-protobuf](https://pkg.go.dev/google.golang.org/protobuf)
- [yalm](https://github.com/go-yaml/yaml/tree/v2.4.0)

## Test commands for the project
These are the commands to run the unit and integration tests of the project

#### Execute All Test
This is the command to run the white box tests, and the test report command
```
go test -v -coverprofile=coverage.out -coverpkg=./... ./...

go tool cover -html=coverage.out
```

This command gets the total coverage of the project
```
go tool cover -func coverage.out
```

## Quick Run Project
Start the terminal to build and run the gRPC server with the HTTP / REST gateway (replace the parameters according to your server):

```
cd cmd/server

go build .

./server -grpc-port=9090 -http-port=8080
```
If we see:

```
2021/02/26 01:26:05 starting gRPC server...
2021/02/26 01:26:05 starting HTTP/REST gateway...
```

## Test Postman
Here is the path of the gRPC api that can be tested with the HTTP / REST endpoint

##### POST

```
localhost:8080/v1/multiply
```
##### Example Json Body
```javascript
{
    "api":"v1",
    "number1":5,
    "number2":2.5
}
 ```

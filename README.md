# WIP PriceFetcher microservice
This is a small microservice with seperation of concerns for logging, metrics and handlers

## Protobuffer installing

### Linux
```
sudo apt install -y protobuf-compiler
```

### MacOS
```
brew install protobuf
```

### GRPC and PRotobuffer package dependencies
```
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go

go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

NOTE: You should add `protoc-gen-go-grpc`to your PATH

```
PATH="${PATH}:$HOME}/go/bin"
```


### Running the service
```
make run
```


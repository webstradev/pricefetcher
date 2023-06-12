build:
	@echo "Building binary..."
	@go build -o bin/pricefetcher

run: build
	@echo "Running binary..."
	@./bin/pricefetcher

proto:
	@echo "Generating protobuf definitions..."
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/service.proto
	@echo "Done generating!"

.PHONY: proto
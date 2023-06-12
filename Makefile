build:
	@echo "Building binary..."
	@go build -o bin/pricefetcher

run: build
	@echo "Running binary..."
	@./bin/pricefetcher
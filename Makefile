.PHONY: build install clean test test-coverage docker-build docker-run

# Binary name
BINARY_NAME=issue-tracker

# Docker image name
DOCKER_IMAGE=cli-task-manager

# Docker volume name
DOCKER_VOLUME=cli-task-manager-data

# Build the application
build:
	go build -o $(BINARY_NAME) ./cmd

# Install the application to $GOPATH/bin
install: build
	go install ./cmd

# Clean build artifacts
clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -f coverage.out

# Run the application
run: build
	./$(BINARY_NAME)

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Build Docker image
docker-build:
	docker build -t $(DOCKER_IMAGE) .

# Create Docker volume if it doesn't exist
docker-volume:
	docker volume create $(DOCKER_VOLUME) || true

# Run Docker container with persistent volume
docker-run: docker-volume
	docker run -it --rm -v $(DOCKER_VOLUME):/root/.cli-task-manager $(DOCKER_IMAGE) $(CMD)

# Build for multiple platforms
build-all:
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-darwin-amd64 ./cmd
	GOOS=darwin GOARCH=arm64 go build -o $(BINARY_NAME)-darwin-arm64 ./cmd
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux-amd64 ./cmd
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME)-windows-amd64.exe ./cmd 
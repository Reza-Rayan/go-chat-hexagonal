APP_NAME=chat-app
MAIN_FILE=cmd/app/main.go
BUILD_DIR=bin
PORT=5050

# Default target
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make run         - Run the app locally"
	@echo "  make build       - Build the binary"
	@echo "  make clean       - Remove binaries"
	@echo "  make tidy        - Clean and verify dependencies"
	@echo "  make test        - Run all tests"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-run  - Run app in Docker"

.PHONY: run
run:
	@echo "🚀 Running $(APP_NAME) on port $(PORT)..."
	go run $(MAIN_FILE)

.PHONY: build
build:
	@echo "🔨 Building binary..."
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

.PHONY: clean
clean:
	@echo "🧹 Cleaning up..."
	rm -rf $(BUILD_DIR)

.PHONY: tidy
tidy:
	@echo "📦 Cleaning and verifying dependencies..."
	go mod tidy
	go mod verify

.PHONY: test
test:
	@echo "🧪 Running tests..."
	go test ./... -v

.PHONY: docker-build
docker-build:
	@echo "🐳 Building Docker image..."
	docker build -t $(APP_NAME):latest .

.PHONY: docker-run
docker-run:
	@echo "🐳 Running app in Docker..."
	docker run -p $(PORT):$(PORT) --env-file .env $(APP_NAME):latest

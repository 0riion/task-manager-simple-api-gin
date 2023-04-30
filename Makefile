APP_NAME=task-manager-api

build:
	@echo "Building $(APP_NAME)..."
	go build -o $(APP_NAME) cmd/main.go

run:
	@echo "Running $(APP_NAME)..."
	go run cmd/main.go

clean:
	@echo "Cleaning $(APP_NAME)..."
	rm -f $(APP_NAME)

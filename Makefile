DSN="host=localhost port=5534 user=postgres password=password dbname=freementors sslmode=disable timezone=UTC connect_timeout=5"
BINARY_NAME=freementors

## build: Build binary
build:
	@echo "Building server..."
	go build -o ${BINARY_NAME} ./cmd/api/
	@echo "Binary built"

## run:  builds and runs the application
run: build
	@echo "Starting the server..."
	@env DSN=${DSN} ./${BINARY_NAME} &
	@echo "Server started!"

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "Cleaned"

## start: an alias to run
start: run

## stop: stops the running application
stop: 
	@echo "Stopping the server..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Server stopped..."

## restart: stops and starts the running application
restart: stop start
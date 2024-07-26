BINARY_NAME=simple_server
OUTPUT_DIR=output

.PHONY: build clean

# ~ make build
build:
	go build -o $(OUTPUT_DIR)/$(BINARY_NAME) .

# ~ make clean
clean:
	rm -f $(OUTPUT_DIR)/$(BINARY_NAME)
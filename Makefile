MAIN_PATH=$(shell pwd)/cmd/
NAME=app
BIN_DIR=$(shell pwd)/bin

build:
	CGO_ENABLED=0 go build -v -o $(BIN_DIR)/$(NAME) $(MAIN_PATH)
run:
	LOG_FILE=file.log COUNT=5 TEMP=32 $(BIN_DIR)/$(NAME)
test:
	go test main.go

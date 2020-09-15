build:
	go build -o ./app

run:
	LOG_FILE=file.log ./app

test:
	go test main.go 
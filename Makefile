build:
	go build -o bin/main main.go 
basic:
	./bin/main
all: build basic
profile-google: build
	./bin/main --profile 10 --url https://www.google.com/
build:
	@go build -o bin/learnBackendApi cmd/main.go 


run: build
	@./bin/learnBackendApi
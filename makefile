build:
	@go build -o bin/learnBackendApi cmd/main.go 

run: build
	@./bin/learnBackendApi

# migration
migration:
	@/home/bishalkoirala/go/bin/migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@, $(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

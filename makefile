# Build the Go application
build:
	@go build -o bin/learnBackendApi cmd/main.go

# Run the Go application
run: build
	@./bin/learnBackendApi

# Create migration files
migration:
	@/home/bishalkoirala/go/bin/migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@, $(MAKECMDGOALS))

# Apply migrations (migrate up)
migrate-up:
	@go run cmd/migrate/main.go up

# Rollback migrations (migrate down)
migrate-down:
	@go run cmd/migrate/main.go down

# Force a migration version (to fix dirty state)
migrate-force:
	@/home/bishalkoirala/go/bin/migrate -database "mysql://root:root@tcp(localhost:3306)/testDb" -path ./cmd/migrate/migrations force 20250323081249

# Apply migration up command (with the migration path defined)
migrate:
	@/home/bishalkoirala/go/bin/migrate -database "mysql://root:root@tcp(localhost:3306)/testDb" -path ./cmd/migrate/migrations up

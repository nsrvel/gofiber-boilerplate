# ==============================================================================
# Golang commands
run:
	echo "Running env local"
	go run cmd/api/main.go

# ==============================================================================
# Docker compose commands

composeup:
	echo "Starting local environment"
	docker compose -f docker-compose.yml up -d
composedown:
	echo "Shutdown local environment"
	docker compose -f docker-compose.yml down

# ==============================================================================
# Go migrate postgresql
migrateup:
	migrate -database postgres://root:secret@localhost:5432/SimpleTransaction?sslmode=disable -path migrations up 1
migratedown:
	migrate -database postgres://root:secret@localhost:5432/SimpleTransaction?sslmode=disable -path migrations down 1

tools:
	./go-tools

.PHONY: run composeup composedown migrateup migratedown tools
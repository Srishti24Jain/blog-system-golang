BINARY=bin/engine

DB := blog
DB_DIR := db
POSTGRES_USER := postgres
POSTGRES_HOST := localhost
POSTGRES_PASSWORD := pwd123
POSTGRES_PORT := 5432


.PHONY: fmt
fmt: ## Format all go files
	goimports -w -local "blog/" .


.PHONY: migrate-prepare
migrate-prepare:
	@echo "Installing golang-migrate"
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: db-migrateup
db-migrateup:
	@-migrate -database 'postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(DB)?sslmode=disable' -path $(DB_DIR)/migrations/ up

.PHONY: db-migratedown
db-migratedown:
	@-migrate -database 'postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(DB)?sslmode=disable' -path $(DB_DIR)/migrations/ down

.PHONY: db-migrateforce
db-migrateforce:
	migrate -database 'postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(DB)?sslmode=disable' -path $(DB_DIR)/migrations/ force 1

.PHONY: repository-gen-prepare
repository-gen-prepare:
	go install github.com/volatiletech/sqlboiler/v4@latest
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3@latest

.PHONY: repository-gen
repository-gen:
	@sqlboiler sqlite3
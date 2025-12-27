# ===== Run =====
run:
	go run ./cmd/api/main.go

# ===== Build =====
build:
	go build -o bin/api ./cmd/api/main.go

# ===== Config =====
include .env.make
export


MIGRATE_PATH = migrations
MIGRATE = migrate -path $(MIGRATE_PATH) -database "$(DB_URL)"

# ===== Commands =====

.PHONY: migrate-up migrate-down migrate-drop migrate-force migrate-version migrate-create

print-env:
	@echo "DB_URL='$(DB_URL)'"

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down 1

migrate-drop:
	$(MIGRATE) drop -f

migrate-force:
	$(MIGRATE) force $(version)

migrate-version:
	$(MIGRATE) version

migrate-create:
	migrate create -ext sql -dir $(MIGRATE_PATH) $(name)

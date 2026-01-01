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

.PHONY: migrate-up migrate-down migrate-drop migrate-force migrate-version migrate-create atlas-diff atlas-apply

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

# ===== NEW: Atlas (Modern) =====
# 1. So sánh GORM code vs DB và sinh file migration
diff:
	atlas migrate diff $(name) \
	  --dir "file://migrations" \
	  --to "external://postgres/gorm" \
	  --dev-url "docker://postgres/15/dev?search_path=public" \
	  --external-schema "postgres/gorm:go run ./cmd/atlas/main.go"

# 2. Áp dụng migration (giống migrate-up nhưng dùng Atlas, hoặc cứ dùng migrate-up cũ cũng được)
apply:
	atlas migrate apply \
	  --dir "file://migrations" \
	  --url "$(DB_URL)"

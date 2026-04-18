DB_USER=postgres
DB_PASSWORD=postgres
DB_HOST=localhost
DB_PORT=5432
DB_NAME=postgres
SSL_MODE=disable

POSTGRESQL_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)

.PHONY: migrate-create
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name

.PHONY: migrate-up
migrate-up:
	migrate -database ${POSTGRESQL_URL} -path migrations up

.PHONY: migrate-down
migrate-down:
	migrate -database ${POSTGRESQL_URL} -path migrations down

.PHONY: migrate-force
migrate-force:
	@read -p "Enter version: " version; \
	migrate -database ${POSTGRESQL_URL} -path migrations force $$version

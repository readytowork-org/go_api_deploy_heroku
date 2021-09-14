include .env

DB_URL="postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

# MIGRATE=migrate -source file://migration -database ${DB_URL} -verbose
MIGRATE=docker-compose exec web migrate -path=migration -database ${DB_URL} -verbose

dev:
	go run main.go

migrate-up:
	$(MIGRATE) up 

migrate-down:
	$(MIGRATE) down 

force:
	@read -p  "Which version do you want to force?" VERSION; \
	$(MIGRATE) force $$VERSION

goto:
	@read -p  "Which version do you want to migrate?" VERSION; \
	$(MIGRATE) goto $$VERSION

migrate:
	@read -p  "What is the name of migration?" NAME; \
	migrate create -ext sql -dir migration  $$NAME
-include .env

.SLENT:

DB_URL=postgresql://$(USER):$(PASSWORD)@$(HOST):$(PORT)/$(DB_NAME)?sslmode=disable

tidy:
	@go mod tidy
	@go mod vendor

run:
	@go run cmd/main.go

migration:
	@migrate create -ext sql -dir ./migrations -seq $(name)
migrateup:
	@migrate -path ./migrations -database "$(DB_URL)" -verbose up

migratedown:
	@migrate -path ./migrations -database "$(DB_URL)" -verbose down

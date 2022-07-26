postgres_init:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

create_db:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

drop_db:
	docker exec -it postgres12 dropdb simple_bank

migrate_up:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: create_db, drop_db, postgres_init, migrate_up, migrate_down, sqlc, test
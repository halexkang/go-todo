init:
	docker run --rm --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d postgres

run:
	docker exec -it postgres psql

createdb:
	docker exec -it postgres createdb --username=root --owner=root go-todo

dropdb:
	docker exec -it postgres dropdb go-todo

migrateup:
	migrate -path db/migrations -database "postgresql://root:postgres@localhost:5432/go-todo?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:postgres@localhost:5432/go-todo?sslmode=disable" -verbose down
.PHONY: init rundb
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root userDB

dropdb:
	docker exec -it postgres12 dropdb userDB

migratecreate:
	migrate create -ext sql -dir db/migration -seq  init_schema

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/userDB?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/userDB?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v "C:\Users\aliki\Desktop\GO\src\user-api:/src" -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migratecreate migrateup migratedown sqlc test
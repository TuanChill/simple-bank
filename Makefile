postgres:
	docker run --name simple-bank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin -d postgres

createdb:
	docker exec -it simple-bank createdb --username=root --owner=root simple-bank

dropdb: 
	docker exec -it simple-bank dropdb simple-bank

migrateup:
	migrate -path db/migration -database "postgresql://root:admin@localhost:5432/simple-bank?sslmode=disable" -verbose up 

migrateup1:
	migrate -path db/migration -database "postgresql://root:admin@localhost:5432/simple-bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:admin@localhost:5432/simple-bank?sslmode=disable" -verbose down 

migratedown1:
	migrate -path db/migration -database "postgresql://root:admin@localhost:5432/simple-bank?sslmode=disable" -verbose down 1

sqlcinit: 
	sqlc init

sqlcgen:
	sqlc generate

test: 
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go -package mockdb  github.com/simple-bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup2 migrateup1 migratedown2 migratedown1 sqlcgen sqlcinit test server mock


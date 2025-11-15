postgres:
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=divymbohra -e POSTGRES_PASSWORD=divymbohra -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=divymbohra --owner=divymbohra

dropdb:
	psql -U divymbohra -c "DROP DATABASE simplebank;"

migrateup: 
	migrate -path db/migration -database "postgresql://divymbohra:divymbohra@localhost:5432/simplebank?sslmode=disable" -verbose up

migrateup1: 
	migrate -path db/migration -database "postgresql://divymbohra:divymbohra@localhost:5432/simplebank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://divymbohra:divymbohra@localhost:5432/simplebank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://divymbohra:divymbohra@localhost:5432/simplebank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate
test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server mock

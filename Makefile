postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=divymbohra -e POSTGRES_PASSWORD=divymbohra -d postgres:12-alpine
createdb:
	docker exec -it postgres12 createdb --username=divymbohra --owner=divymbohra

dropdb:
	psql -U divymbohra -c "DROP DATABASE simplebank;"

migrateup: 
	migrate -path db/migration -database "postgresql://divymbohra:divymbohra@localhost:5432/simplebank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://divymbohra:divymbohra@localhost:5432/simplebank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test

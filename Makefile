DB_URL=postgresql://root:secretpass@localhost:5432/bank?sslmode=disable
postgres:
	docker run --name postgres01 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secretpass -d postgres:15.2-alpine
createdb:
	docker exec -it postgres01 createdb --username=root --owner=root bank
dropdb:
	docker exec -it postgres01 dropdb --username=root --owner=root bank
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1
migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store
db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock migrateup1 migratedown1 db_docs db_schema
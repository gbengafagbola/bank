postgres:
	docker run --name postgres01 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secretpass -d postgres:15.2-alpine
createdb:
	docker exec -it postgres01 createdb --username=root --owner=root bank
dropdb:
	docker exec -it postgres01 dropdb --username=root --owner=root bank
migrateup:
	migrate -path db/migration -database "postgresql://root:secretpass@localhost:5432/bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secretpass@localhost:5432/bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store
    
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock
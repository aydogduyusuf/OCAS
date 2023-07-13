postgres:
	sudo docker run --name OCASDB -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=group33ocas -d postgres:14-alpine
createdb:
	sudo docker exec -it OCASDB createdb --username=root --owner=root OCASDB
dropdb:
	sudo docker exec -it OCASDB dropdb OCASDB
migrateinit:
	migrate create -ext sql -dir src/DB/migration -seq init_schema
migrateup: 
	migrate -path src/DB/migration -database "postgresql://root:group33ocas@localhost:5432/OCASDB?sslmode=disable" -verbose up
migratedown:
	migrate -path src/DB/migration -database "postgresql://root:group33ocas@localhost:5432/OCASDB?sslmode=disable" -verbose down
migrateevents:
	migrate create -ext sql -dir src/DB/migration -seq events
migrateup1:
	migrate -path src/DB/migration -database "postgresql://root:group33ocas@localhost:5432/OCASDB?sslmode=disable" -verbose up 1
migratedown1:
	migrate -path src/DB/migration -database "postgresql://root:group33ocas@localhost:5432/OCASDB?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
.PHONY: postgres createdb dropdb migrateinit migrateup migratedown sqlc migrateevents migrateup1 migratedown1
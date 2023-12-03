postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=167916 -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root enigma_laundry_docker

dropdb:
	docker exec -it postgres16 dropdb enigma_laundry_docker

migrateup:
	migrate -path db/migration -database "postgresql://root:167916@localhost:9090/enigma_laundry_docker?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:167916@localhost:9090/enigma_laundry_docker?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown
rabbitmq:
	docker start rabbitmq-container

redis:
	docker start redis-container

postgres:
	docker run --name pg16 -e POSTGRES_USER=root  -e POSTGRES_PASSWORD=devs -p:5432:5432 -d postgres:16-alpine

createdb:
	docker exec -it pg16 createdb --username=root --owner=root leaderboard

dropdb:
	docker exec -it pg16 dropdb leaderboard

migrateup:
	migrate -path db/migration -database "postgresql://root:devs@localhost:5432/leaderboard?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:devs@localhost:5432/leaderboard?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: rabbitmq redis postgres createdb dropdb migrateup migratedown sqlc


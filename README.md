# leaderboard
A simple leaderboard application that consumes events from rabbitmq and updates the leaderboard.

Stack: RabbitMQ, Redis, Postgres
## Setup
### RabbitMQ
RabbitMQ

docker run -d --name rabbitmq-container -p 5672:5672 -p 15672:15672 rabbitmq:3-management

This command creates a Docker container named "rabbitmq-container" based on the "rabbitmq:3-management" image. The -p flag maps the RabbitMQ ports 5672 (AMQP) and 15672 (RabbitMQ management UI) to your host machine.

Start
	docker start rabbitmq-container
Stop
	docker stop rabbitmq-container
Remove
	docker rm rabbitmq-container


You can access the RabbitMQ Management UI by opening a web browser and navigating to http://localhost:15672/. 


### Redis
Pull latest redis
	docker pull redis

Create redis container named redis-container
	docker run -d --name redis-container -p 6379:6379 redis 
	The -p flag maps the redis port 6379 to your host machine port 6379

Start container 

	docker start redis-container

Stop container

	docker stop redis-container


For connecting to container

	docker exec -it redis-container bash

Access redis cli

	redis-cli

### postgres
Setup postgres
 docker pull postgres:16-alpine
 --images with alpine tags are smaller in size


Start a container
	docker run --name some-postgres-name -e POSTGRES_USER=root  -e POSTGRES_PASSWORD=devs -p:5432:5432 -d postgres:16-alpine


psql
	docker exec -it pg16 psql -U root
	\q to exit

## How to run 

make postgres
make createdb
make migrateup
go build
go run *.go

Go to RabbitMQ management UI, go to leaderboard queue and publish a message.

Sample - {"user_name": "Abhi", "user_id": 1, "points":201}

Voila!



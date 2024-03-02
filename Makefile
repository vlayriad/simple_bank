POSTGRES_USER = root
POSTGRES_NAME = simple_bank
POSTGRES_PASS = secret
DB_HOST 	  = localhost
DB_PATH 	  = database/migration
DB_PORT 	  = 5432
DB_DRIVER 	  = postgresql
DB_SOURCE 	  = postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable
DB_SSL_MODE   = disable

DOCKER_IMAGES 		= postgres:16-alpine
DOCKER_PORT 		= 5432
DOCKER_CONTAINER 	= postgres16

# IF INSTALLED DOCKER & WSL IN YOUR WINDOWS
# PLEASE DELETE wsl BEFORE COMMAND
# wsl docker run --name to -> docker run --name
postgres:
	wsl docker run --name $(DOCKER_CONTAINER) -p $(DB_PORT):$(DOCKER_PORT) -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASS) -d $(DOCKER_IMAGES)

createdb:
	wsl docker exec -it $(DOCKER_CONTAINER) createdb --username=$(POSTGRES_USER) --owner=$(POSTGRES_USER) $(POSTGRES_NAME)

dropdb:
	wsl docker exec -it $(DOCKER_CONTAINER) dropdb $(POSTGRES_NAME)

migrateup:
	migrate -path $(DB_PATH) -database "$(DB_DRIVER)://$(POSTGRES_USER):$(POSTGRES_PASS)@$(DB_HOST):$(DOCKER_PORT)/$(POSTGRES_NAME)?sslmode=$(DB_SSL_MODE)" -verbose up

migratedown:
	echo y | migrate -path $(DB_PATH) -database "$(DB_DRIVER)://$(POSTGRES_USER):$(POSTGRES_PASS)@$(DB_HOST):$(DOCKER_PORT)/$(POSTGRES_NAME)?sslmode=$(DB_SSL_MODE)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres dropdb createdb migrateup migratedown sqlc
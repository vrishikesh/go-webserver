up:
	docker compose up

build:
	CGO_ENABLED=0 GOOS=linux go build -v -o ./dist/main .

up_build: down build up

down:
	docker compose down

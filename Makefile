build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

seed:
	docker-compose run backend 'go run /app/src/commands/populateUsers.go'
build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

seed:
	go run /app/src/commands/populateUsers.go
	go run /app/src/commands/populateProducts.go
	go run /app/src/commands/populateOrders.go

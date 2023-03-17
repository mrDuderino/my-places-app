.PHONY: migrate-create migrate-up migrate-down run-db connect-db

migrate-create:
	migrate create -ext sql -dir ./schema -seq init

migrate-up:
	migrate -path ./schema -database 'postgres://postgres:qwertysha@localhost:5433/postgres?sslmode=disable' -verbose up

migrate-down:
	migrate -path ./schema -database 'postgres://postgres:qwertysha@localhost:5433/postgres?sslmode=disable' -verbose down

run-db:
	docker run --name=my-places-db -e POSTGRES_PASSWORD="qwertysha" -p 5433:5432 -d postgres

connect-db:
	docker exec -it my-places-db /bin/bash
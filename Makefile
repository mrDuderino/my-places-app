.PHONY: migrate create
migrate create:
	migrate create -ext sql -dir ./schema -seq init

.PHONY: migrate run
migrate run:
	migrate -path ./schema -database 'postgres://postgres:qwertysha@localhost:5433/postgres?sslmode=disable' up

.PHONY: postgres connect
postgres connect:
	docker exec -it my-places-db /bin/bash


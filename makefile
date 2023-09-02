run:
	go run cmd/app/main.go

run_docker:
	docker-compose build && docker-compose up

make migration:
	migrate create -ext sql -dir ./schema -seq init

migrate:
	migrate -path ./schema -database 'root:$(PASSWORD)@tcp(127.0.0.1:3306)/db'

rollback:
	migrate -path ./schema -database 'root:$(PASSWORD)@tcp(127.0.0.1:3306)/db'

path_reset:
	export PATH=$PATH:$(go env GOPATH)/bin
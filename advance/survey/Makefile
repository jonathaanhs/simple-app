
test:
	go test ./...

run:
	docker-compose up -d

stop:
	docker-compose down

migrate-up:
	migrate -database "postgresql://survey:survey@:5432/survey?sslmode=disable" -path ./sql up

migrate-down:
	migrate -database "postgresql://survey:survey@:5432/survey?sslmode=disable" -path ./sql down
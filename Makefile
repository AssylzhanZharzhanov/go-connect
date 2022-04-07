.SILENT:

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/api/main.go

run: build
	docker-compose up -d --remove-orphans --build server

docker-build-local: run
	docker-compose up -d

test:
	go test -cover ./...

stop:
	docker stop server && docker rm server

rebuild: stop
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/api/main.go && docker-compose up -d --remove-orphans --build server

migrate-up:
	migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up

migrate-down:
	migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" down

migrate-force:
	migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" force 1

migrate-drop:
	migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" drop


.DEFAULT_GOAL := run
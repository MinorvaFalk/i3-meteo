run-api:
	go run cmd/api/main.go

run-worker:
	go run cmd/worker/main.go

migrate-up:
	go run cmd/migrate/*.go -dir=migrations postgres postgresql://root:root@localhost:5432/database up

migrate-down:
	go run cmd/migrate/*.go -dir=migrations postgres postgresql://root:root@localhost:5432/database down

postgres:
	docker run -itd \
		--name postgres \
		-p 5432:5432 \
		-e POSTGRES_USER=root \
		-e POSTGRES_PASSWORD=root \
		-e POSTGRES_DB=database \
		postgres:alpine

redis:
	docker run -itd \
		--name redis \
		-p 6379:6379 \
		redis:alpine

build:
	go mod tidy &&
	docker build -f ./infra/Dockerfile -t i3-meteo:0.0.0 .

compose:
	docker compose -p i3-meteo -f infra/docker-compose.yml up -d
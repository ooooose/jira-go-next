build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down

ps:
	docker compose ps

logs:
	docker compose logs -f go

lint:
	docker compose run --rm go golint .

fmt:
	docker compose run --rm go gofmt -w .

shell:
	docker compose run --rm go bash


.PHONY: start
start: down up fixtures logs

.PHONY: prepare
prepare:
	@if [ ! -f .env ]; then cp .env.example .env; fi

.PHONY: up
up: prepare
	@docker compose up -d

.PHONY: down
down: prepare
	@docker compose down

.PHONY: fixtures
fixtures: prepare
	@docker compose exec db mongorestore \
		--uri mongodb://db:27017 \
		--gzip /backup/cinema

.PHONY: logs
logs: prepare
	@docker compose logs -f

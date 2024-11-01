.PHONY: up
up:
	@docker compose up -d

.PHONY: down
down:
	@docker compose down

.PHONY: fixtures
fixtures:
	@docker compose exec db mongorestore \
		--uri mongodb://db:27017 \
		--gzip /backup/cinema

.PHONY: logs
logs:
	@docker compose logs -f

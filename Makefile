DEV_COMPOSE = docker/docker-compose.dev.yaml
PROD_COMPOSE = docker/docker-compose.prod.yaml

.PHONY: dev-up dev-down dev-restart dev-logs
.PHONY: prod-up prod-down prod-restart prod-logs
.PHONY: migrate

## ---------------------------
## Development
## ---------------------------

dev-up:
	docker compose -f $(DEV_COMPOSE) --env-file .env up -d

dev-down:
	docker compose -f $(DEV_COMPOSE) --env-file .env down -v

dev-restart: dev-down dev-up

dev-logs:
	docker compose -f $(DEV_COMPOSE) logs -f db

## ---------------------------
## Production
## ---------------------------

prod-up:
	docker compose -f $(PROD_COMPOSE) up -d --build

prod-down:
	docker compose -f $(PROD_COMPOSE) down -v

prod-restart: prod-down prod-up

prod-logs:
	docker compose -f $(PROD_COMPOSE) logs -f

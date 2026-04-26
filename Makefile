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

VENV=venv
PYTHON=$(VENV)/bin/python
PIP=$(VENV)/bin/pip
USERS_CSV_FILE=tools/generate_users/users.csv

# Если нет venv — создаём
$(VENV)/bin/activate: 
	python3 -m venv $(VENV)
	$(PIP) install --upgrade pip
	$(PIP) install -r requirements.txt

# Команда для генерации пользователей
users-generate: $(VENV)/bin/activate
	$(PYTHON) tools/generate_users/generate_users.py --count $(COUNT) --output $(OUTPUT)

# Значения по умолчанию
export COUNT ?= 1000000
export OUTPUT ?= $(USERS_CSV_FILE)

DB_CONTAINER=social_db_dev
DB_USER=postgres
DB_NAME=social
CSV_FILE=users.csv

users-load:
	@if [ ! -f $(USERS_CSV_FILE) ]; then echo "CSV file '$(USERS_CSV_FILE)' not found! Run make generate first."; exit 1; fi
	@echo "[*] Copying $(USERS_CSV_FILE) into container $(DB_CONTAINER)..."
	docker cp $(USERS_CSV_FILE) $(DB_CONTAINER):/$(CSV_FILE)
	@echo "[*] Running COPY inside PostgreSQL..."
	docker exec -i $(DB_CONTAINER) psql -U $(DB_USER) -d $(DB_NAME) -c " \
		COPY users(first_name, second_name, birthdate, biography, city, password_hash) \
		FROM '/$(CSV_FILE)' DELIMITER ',' CSV;"
	@echo "[✓] Data loaded into PostgreSQL!"
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

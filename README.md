# OTUS Social Network

Simple social network backend written in Go. The service exposes a small HTTP API for user registration, user lookup, and login with JWT token generation. PostgreSQL is used as the main data store.

## Stack

- Go `1.25.5`
- PostgreSQL `15`
- Docker and Docker Compose
- Chi router
- JWT auth

## Project Structure

```text
.
├── cmd/app                  # application entrypoint
├── docker                   # docker compose files
├── internal
│   ├── auth                 # login and JWT helpers
│   ├── config               # env-based configuration
│   ├── db                   # postgres connection
│   ├── http                 # router and handlers
│   └── user                 # user domain logic
├── migrations               # database bootstrap SQL
├── Dockerfile
└── Makefile
```

## Environment Variables

The application reads variables from `.env` if the file exists, otherwise from system environment variables.

| Variable | Default | Description |
| --- | --- | --- |
| `APP_PORT` | `8080` | HTTP server port |
| `DB_HOST` | `localhost` | PostgreSQL host |
| `DB_PORT` | `5432` | PostgreSQL port |
| `DB_USER` | `postgres` | PostgreSQL user |
| `DB_PASSWORD` | `postgres` | PostgreSQL password |
| `DB_NAME` | `social` | PostgreSQL database name |

Example `.env`:

```env
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=social
```

## Run In Development

Start PostgreSQL and pgAdmin:

```bash
make dev-up
```

Useful development commands:

```bash
make dev-logs
make dev-down
make dev-restart
```

Run the Go application locally:

```bash
go run ./cmd/app
```

Development services:

- PostgreSQL: `localhost:5432`
- pgAdmin: `http://localhost:5050`
- API: `http://localhost:8080`

## Run In Production Mode

Build and start the app together with PostgreSQL:

```bash
make prod-up
```

Useful production commands:

```bash
make prod-logs
make prod-down
make prod-restart
```

The production compose setup publishes the API on `http://localhost:8080`.

## Database

On container startup, SQL files from `migrations/001_init.sql` are mounted into the PostgreSQL initialization directory.

Current schema includes a single `users` table with:

- `id` UUID primary key
- `first_name`
- `second_name`
- `birthdate`
- `biography`
- `city`
- `password_hash`

## HTTP API

### `POST /user/register`

Creates a user.

Request example:

```bash
curl -X POST http://localhost:8080/user/register \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Ivan",
    "second_name": "Ivanov",
    "birthdate": "1990-01-01",
    "biography": "Backend developer",
    "city": "Moscow",
    "password": "secret123"
  }'
```

Response example:

```json
{
  "user_id": "uuid"
}
```

### `GET /user/get/{id}`

Returns a user by id.

Request example:

```bash
curl http://localhost:8080/user/get/<user_id>
```

### `POST /login`

Checks user credentials and returns a JWT token.

Request example:

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "id": "<user_id>",
    "password": "secret123"
  }'
```

Response example:

```json
{
  "token": "jwt-token"
}
```

## Notes

- Passwords are stored as SHA-256 hashes.
- JWT signing key is currently hardcoded in the source and should be moved to configuration before real production use.
- There are no automated tests in the repository yet.

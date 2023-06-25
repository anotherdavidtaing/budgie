LOCAL_DATABASE_URL="postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
MIGRATIONS_PATH="./backend/db/migrations"

is-local-env:
	@if [ "$(ENVIRONMENT)" != "local" ]; then \
		echo "This command is not available in this environment: $(ENVIRONMENT)"; \
		exit 1; \
	fi

start-dev-db:
	@$(MAKE) is-local-env
	@echo "Setting up local development database"
	@cd ./backend && docker compose --profile development up -d
	@echo "Running database migrations"
	@migrate -database $(LOCAL_DATABASE_URL) -path $(MIGRATIONS_PATH) up
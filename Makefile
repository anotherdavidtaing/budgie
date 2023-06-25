start-dev-db:
	@if [ "$(ENVIRONMENT)" = "local" ]; then \
		echo "Setting up local development database"; \
		cd ./backend && docker compose --profile development up -d; \
	else \
		echo "This command is not available in this environment: $(ENVIRONMENT)"; \
		exit 1; \
	fi
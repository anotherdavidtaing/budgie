start-dev-db:
	@echo "Setting up local development database"
	cd ./backend && docker compose --profile development up -d

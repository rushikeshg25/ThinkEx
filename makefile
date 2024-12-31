dev-server:
	@echo "Starting development server"
	@cd server && make run

dev-web:
	@echo "Starting development web"
	@cd web && bun run dev

docker-run:
	@echo "Starting docker containers"
	@docker compose up

.PHONY: dev-server dev-web
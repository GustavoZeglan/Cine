up:
	@echo "Starting Docker Containers..."
	docker compose -f compose.yml up -d
	@echo "Docker containers are started!"

down:
	@echo "Finishing Docker Containers..."
	docker compose -f compose.yml down
	@echo "Done!"

stop:
	@echo "Stopping Docker Containers..."
	docker compose -f compose.yml stop
	@echo "Stopped!"
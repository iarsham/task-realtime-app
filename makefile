# Target to run tests in user-service
test-user:
	cd user-service && \
	go test ./repository

# Target to run tests in chat-service
test-chat:
	cd chat-service && \
	go test ./repository

# Target to run tests in all services
test:
	make test-user
	make test-chat

# Target to run dev environment
run-dev:
	docker-compose -f docker-compose-dev.yaml up --build -d

# Target to stop dev environment
stop-dev:
	docker-compose -f docker-compose-dev.yaml down

# Target to run production environment
run-prod:
	docker-compose up --build -d

# Target to stop production environment
stop-prod:
	docker-compose down

# Target to run logs in user-service
log-user:
	docker-compose logs -f user-service

# Target to run logs in chat-service
log-chat:
	docker-compose logs -f chat-service

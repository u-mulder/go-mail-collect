# Variables
DOCKER=docker
DOCKER_COMPOSE=docker-compose
GO_CONTAINER=go

# start dev-service
.PHONY: start
start: 
	@echo "==> Startng go service" 
	@${DOCKER_COMPOSE} up -d

.PHONY: up
up: start 

# stop dev service
.PHONY: stop
stop: 
	@echo "==> Stopping go service" 
	@${DOCKER_COMPOSE} down

.PHONY: down
down: stop 

# restart dev service
.PHONY: restart
restart: stop start

# rebuild go-container
.PHONY: rebuildgo
rebuildgo: 
	@echo "==> Stopping go service" 
	@${DOCKER_COMPOSE} up -d --no-deps --build ${GO_CONTAINER}

# Logs of go container
.PHONY: logs
logs: 
	@echo "==> Logs of go service" 
	@${DOCKER} logs mailcollector_go_1

# List of containers
.PHONY: contls
contls: 
	@echo "==> List running containers" 
	@${DOCKER} container ls

# More commands can be included here
-include local.mk
DC := docker-compose -f ./docker-compose.yaml

.PHONY: build
build:
	${DC} build

.PHONY: up
up:
	${DC} up -d

.PHONY: down
down:
	${DC} down

.PHONY: clean 
clean:
	docker volume prune -f

.PHONY: restart
restart: down clean build up

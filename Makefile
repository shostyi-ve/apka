
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

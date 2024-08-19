ifndef $(GOROOT)
    GOROOT=$(shell go env GOROOT)
    export GOROOT
endif

include .env
export

DIR=$(notdir $(shell pwd))
export DIR

dev:
	@echo "\n\t🧬\n"
	docker-compose -p $(DIR) up --remove-orphans

.PHONY: dev
# get DB_URL from .env
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

up:
	cd sql/schema && goose postgres "$(DB_URL)" up
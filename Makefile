ifneq (,$(wildcard config/.env))
    include config/.env
endif

create_env:
	cp config/.env.example config/.env

build:
	@docker build -t rndserver .

start_with_logs:
	@docker compose -f docker-compose.yml --env-file config/.env up

start:
	@docker compose -f docker-compose.yml --env-file config/.env up -d
	
stop:
	@docker compose -f docker-compose.yml  --env-file config/.env stop

kill:
	@docker compose -f docker-compose.yml  --env-file config/.env kill

down:
	@docker compose -f docker-compose.yml  --env-file config/.env down --rmi local

restart: down start

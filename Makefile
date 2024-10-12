ifneq ($(shell docker compose version 2>/dev/null),)
	DC=docker compose
else ifneq ($(shell docker-compose --version 2>/dev/null),)
	DC=docker-compose
else
	$(error ************  docker compose or docker-compose not found. ************)
endif

up:
	@echo "Start BTCGO Project"
	$(DC) -f docker-compose.yml up -d --build
down:
	@echo "Stop BTCGO Project"
	$(DC) -f docker-compose.yml down --remove-orphans

restart:
	@echo "Restaring BTCGO Project"
	make down && make up

update_version:
	@echo "Updating BTCGO Version"
	git fetch && git pull
	make restart

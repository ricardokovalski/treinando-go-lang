
check-env:
ifeq ($(shell [ -f .env ] && echo -n yes),)
	@echo " -------------------------------------------\nyou need to create your `.env` file\n-------------------------------------------"
	exit 1
endif

run: check-env
	@docker compose up
	
run-debug: check-env
	DEBUG="true" docker-compose up 

bash: check-env
	@docker compose run -it --rm --entrypoint=sh go-dev

vendor: 
	@docker compose run -it --rm --entrypoint=go go-dev mod vendor

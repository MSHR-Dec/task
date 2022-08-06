.PHONY: init
init: start/mysql build/migration migrate
start/mysql:
	docker compose up -d mysql
build/migration:
	docker build -f docker/migration/Dockerfile -t task-migration .
migrate:
	docker run --rm --network task_task task-migration

.PHONY: run
run: run/store run/app
run/store:
	docker compose up -d mysql redis
run/app:
	docker compose up -d backend frontend

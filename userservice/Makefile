createdb:
	docker compose exec -it postgres createdb --username=postgres  --owner=postgres telegrammanager

dropdb:
	docker compose exec -it postgres dropdb --username=postgres telegrammanager

migrateup:
	sudo migrate -path db/migration -database "postgres://postgres:mypassword@localhost:6543/telegrammanager?sslmode=disable" -verbose up

migratedown:
	sudo migrate -path db/migration -database "postgres://postgres:mypassword@localhost:6543/telegrammanager?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

run:
	go run cmd/telegrammanager/main.go	

dcu:
	docker compose up -d

dcub:
	docker compose up -d --build

dcd:
	docker compose down

.PHONY: createdb dropdb migrateup migratedown test run dcup dcdown dcub
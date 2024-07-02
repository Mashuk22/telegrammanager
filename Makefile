createdb:
	docker compose exec -it postgres createdb --username=postgres  --owner=postgres telegrammanager

dropdb:
	docker compose exec -it postgres dropdb --username=postgres telegrammanager

migrateup:
	./migrate -path db/migration -database "postgresql://postgres:mypassword@localhost:6543/telegrammanager?sslmode=disable" -verbose up

migratedown:
	./migrate -path db/migration -database "postgresql://postgres:mypassword@localhost:6543/telegrammanager?sslmode=disable" -verbose down

.PHONY: createdb dropdb migrateup migratedown
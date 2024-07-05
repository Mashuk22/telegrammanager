uscreatedb:
	docker compose exec -it postgres createdb --username=postgres  --owner=postgres telegrammanager

usdropdb:
	docker compose exec -it postgres dropdb --username=postgres telegrammanager

usmigrateup:
	sudo migrate -path userservice/db/migration -database "postgres://postgres:mypassword@localhost:6543/telegrammanager?sslmode=disable" -verbose up

usmigratedown:
	sudo migrate -path userservice/db/migration -database "postgres://postgres:mypassword@localhost:6543/telegrammanager?sslmode=disable" -verbose down

ustest:
	go test -v -cover userservice/...

usrun:
	go run userservice/cmd/telegrammanager/main.go	


.PHONY: uscreatedb usdropdb usmigrateup usmigratedown ustest usrun
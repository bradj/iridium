.PHONY: iridium run clean db-start db-generate db-seed

APP := "iridium"

iridium: clean
	go build

run:
	./${APP}

clean:
	rm -f ${APP}

db-start:
	docker run -itd --rm -p 5432:5432 -e POSTGRES_USER="iridium" -e POSTGRES_PASSWORD="123456789" postgres

db-generate:
	go generate

db-migrate:
	sql-migrate up

.SILENT:

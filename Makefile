APP := "iridium"

iridium: clean
	go build

run:
	./${APP}

.PHONY: clean
clean:
	rm -f ${APP}

.PHONY: db-start
db-start:
	docker run -itd --rm -p 5432:5432 -e POSTGRES_USER="iridium" -e POSTGRES_PASSWORD="123456789" postgres
	
.PHONY: db-generate
db-generate:
	sqlboiler --wipe psql

.SILENT:
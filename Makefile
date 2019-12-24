APP := "iridium"

iridium: clean
	go build

run:
	./${APP}

.PHONY: clean
clean:
	rm ${APP}

.PHONY: start-db
start-db:
	docker run -itd --rm -p 5432:5432 -e POSTGRES_USER="iridium" -e POSTGRES_PASSWORD="123456789" postgres
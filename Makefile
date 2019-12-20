APP := "iridium"

iridium: clean
	go build

run:
	./${APP}

.PHONY: clean
clean:
	rm ${APP}
.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go
run: build
	sudo ./.bin/bot
build-container:
	sudo docker build --tag telegram-pocket-bot:v0.1 .
start-container:
	sudo docker run --name telegram-pocket-bot -p 80:80 --env-file .env telegram-pocket-bot:v0.1
prepare-application:
	cat ./samples/.env > .env

run:
	go run ./cmd/url-shortener
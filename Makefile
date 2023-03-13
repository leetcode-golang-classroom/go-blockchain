build:
	@go build -o ./bin/blocker

run: build
	@./bin/blocker
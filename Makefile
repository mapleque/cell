.PHONY: run, build

build:
	-mkdir bin
	go build -o bin/server main/server.go

run:
	cd main && go run server.go


.PHONY: test-server test-html-client run-server-dev run-html-client-dev run install release help

help:
	@echo 'Thanks for your contribution, make can help you more!'
	@echo 'Command list:'
	@echo '	make run-server-dev'
	@echo '		run the server in development mode'
	@echo '	make run-html-client-dev'
	@echo '		run the html client in development mode'
	@echo '	make install'
	@echo '		build both server and client into bin'
	@echo '	make release'
	@echo '		build release packages for multiple platform'

test-server:
	echo 'TODO'

test-html-client:
	echo 'TODO'

run-server-dev: main/.env bin/www
	cd main && go run server.go

run-html-client-dev: client/html/node_modules
	cd client/html && npm run start

run: bin/server bin/.env bin/www bin/templates
	cd bin && chmod +x server && ./server

install: bin/server bin/.env bin/www bin/templates
	@echo 'Install success, please setup env (or edit bin/.env) for config and run.'

release: releases/cell-mac.zip releases/cell-linux.zip releases/cell-win64.zip releases/cell-win32.zip
	@echo 'Build release packages success:'
	@echo '	releases/cell-mac.zip'
	@echo '	releases/cell-linux.zip'
	@echo '	releases/cell-win32.zip'
	@echo '	releases/cell-win64.zip'

releases/cell-mac.zip: releases/cell-mac bin/www bin/templates releases/.env
	zip -r releases/cell-mac.zip releases/cell-mac bin/templates bin/www releases/.env

releases/cell-linux.zip: releases/cell-linux bin/www bin/templates releases/.env
	zip -r releases/cell-linux.zip releases/cell-linux bin/templates bin/www releases/.env

releases/cell-win64.zip: releases/cell-win64 bin/www bin/templates releases/.env
	zip -r releases/cell-win64.zip releases/cell-win64 bin/templates bin/www releases/.env

releases/cell-win32.zip: releases/cell-win32 bin/www bin/templates releases/.env
	zip -r releases/cell-win32.zip releases/cell-win32 bin/templates bin/www releases/.env

releases/cell-mac: releases
	GOOS=darwin GOARCH=amd64 go build -o releases/cell-mac main/server.go

releases/cell-linux: releases
	GOOS=linux GOARCH=amd64 go build -o releases/cell-linux main/server.go

releases/cell-win64: releases
	GOOS=windows GOARCH=amd64 go build -o releases/cell-win64.exe main/server.go

releases/cell-win32: releases
	GOOS=windows GOARCH=386 go build -o releases/cell-win32.exe main/server.go

releases/.env: releases
	cp -r .env.example releases/.env

releases:
	-mkdir releases

main/.env:
	cp .env.example main/.env

client/html/node_modules:
#	cd client/html && npm install
	-mkdir client/html/node_modules

bin/.env: bin
	cp .env.example bin/.env

bin/server: bin
	go build -o bin/server main/server.go

bin/www: bin client/html/build
	cp -r client/html/build bin/www

bin/templates: bin
	cp -r templates bin/

bin:
	mkdir bin

client/html/build: client/html/node_modules
#	cd client/html && npm run build
	-mkdir client/html/build

clean:
	-rm -rf bin
	-rm -rf releases
	-rm -rf client/html/build

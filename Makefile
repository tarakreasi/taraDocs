.PHONY: run build build-assets build-linux build-windows release dist clean

run:
	npm run start

build-assets:
	npm run build

build-linux: build-assets
	GOOS=linux GOARCH=amd64 go build -o bin/taradocs-linux cmd/server/main.go

build-windows: build-assets
	GOOS=windows GOARCH=amd64 go build -o bin/taradocs-windows.exe cmd/server/main.go

release: build-linux build-windows

dist: release
	mkdir -p dist/linux
	mkdir -p dist/windows
	
	# Linux Bundle
	cp bin/taradocs-linux dist/linux/server
	cp -r views dist/linux/views
	cp -r public dist/linux/public
	cp .env.example dist/linux/.env.example
	
	# Windows Bundle
	cp bin/taradocs-windows.exe dist/windows/server.exe
	cp -r views dist/windows/views
	cp -r public dist/windows/public
	cp .env.example dist/windows/.env.example
	
	@echo "Distribution created in dist/"

clean:
	rm -rf bin public/build dist

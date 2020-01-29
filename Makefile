PREFIX=/usr/local

build: format
	go build

format:
	gofmt -s -w .

install: build
	mkdir -p $(PREFIX)/bin
	cp -f agg $(PREFIX)/bin

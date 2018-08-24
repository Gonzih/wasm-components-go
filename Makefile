GOVERSION:=go1.11rc2
WASM_HELPERS:=wasm_exec.html wasm_exec.js

all: clean go $(WASM_HELPERS) test.wasm

%.wasm:
	env GO111MODULE=on GOROOT=$(HOME)/go-projects/src/wasm-test/go/ GOARCH=wasm GOOS=js $(GOVERSION) build -o $@ *.go

wasm_exec.%:
	cp go/misc/wasm/$@ .

clean:
	rm -f test.wasm

godoc:
	env GO111MODULE=on GOROOT=$(HOME)/go-projects/src/wasm-test/go/ godoc -http=:6060

server-main:
	env GO111MODULE=on $(GOVERSION) build -o server-main server/main.go

run-server: server-main
	@echo http://localhost:3000/wasm_exec.html
	./server-main

setup:
	go get golang.org/dl/$(GOVERSION)
	$(GOVERSION) download

go:
	wget https://dl.google.com/go/$(GOVERSION).src.tar.gz -O /tmp/$(GOVERSION).src.tar.gz
	tar xvzf /tmp/$(GOVERSION).src.tar.gz

.PHONY: setup clean godoc run-server

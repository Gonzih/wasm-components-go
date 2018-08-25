GOVERSION:=go1.11rc2
WASM_HELPERS:=wasm_exec.html wasm_exec.js

all: clean go setup $(WASM_HELPERS) test.wasm server-main

setup: $(GOPATH)/bin/$(GOVERSION)

%.wasm:
	env GO111MODULE=on GOROOT=$(shell pwd)/go/ GOARCH=wasm GOOS=js $(GOVERSION) build -o $@

test:
	$(GOVERSION) test

wasm_exec.%:
	cp go/misc/wasm/$@ .

clean:
	rm -f test.wasm

godoc:
	env GO111MODULE=on GOROOT=$(shell pwd)/go/ godoc -http=:6060

server-main:
	env GO111MODULE=on $(GOVERSION) build -o server-main server/main.go

run-server: server-main
	@echo http://localhost:3000/wasm_exec.html
	./server-main

$(GOPATH)/bin/$(GOVERSION):
	go get golang.org/dl/$(GOVERSION)
	$(GOVERSION) download

go:
	wget https://dl.google.com/go/$(GOVERSION).src.tar.gz -O /tmp/$(GOVERSION).src.tar.gz
	tar xzf /tmp/$(GOVERSION).src.tar.gz

.PHONY: setup clean godoc run-server

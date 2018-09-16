GOVERSION:=go
WASM_HELPERS:=wasm_exec.html wasm_exec.js

all: clean $(WASM_HELPERS) test.wasm server-main

%.wasm:
	env GO111MODULE=on GOARCH=wasm GOOS=js $(GOVERSION) build -o $@

test:
	env GO111MODULE=on $(GOVERSION) test

update:
	env GO111MODULE=on $(GOVERSION) get -u

wasm_exec.%:
	curl https://raw.githubusercontent.com/golang/go/go1.11/misc/wasm/$@ > $@

clean:
	rm -f test.wasm

godoc:
	env GO111MODULE=on godoc -http=:6060

server-main:
	env GO111MODULE=on $(GOVERSION) build -o server-main server/main.go

run-server: server-main
	@echo http://localhost:3000/wasm_exec.html
	./server-main

.PHONY: setup clean godoc run-server

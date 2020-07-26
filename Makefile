# always rebuild these targets
.PHONY: test clean

# g++ with all my fancy options
GCC := g++ -DLINUX -m32
GO := CGO_ENABLED=1 GOARCH=386 go

default: build

go-hello.so: go/main.go
	$(GO) build -o build/$@ -buildmode=c-shared $^

.ONESHELL:
update:
	git submodule update --init --recursive

build:
	make go-hello.so

clean:
	rm -rf go-hello.so
	rm -rf go-hello.h

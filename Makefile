CONTENT_SRCS = $(wildcard ) main.go
APPS = chainsys
BLDDIR = build
export GOPATH = $(shell pwd)

all:
	@mkdir -p $(dir $@)
	cd ./src/chainseye.com; go build -o ../../build/bin/${APPS} ${CONTENT_SRCS}
	cp -r ./src/chainseye.com/config ./build
	cp -r ./src/chainseye.com/static ./build

clean:
	rm -fr $(BLDDIR)

.PHONY: clean all

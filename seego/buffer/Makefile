.PHONY: build run all

all: build run

run:
	LD_LIBRARY_PATH=`pwd`/lib ./cgo-buffer

build:
	gcc -fpic -shared -I./include ./include/*.cc -o lib/libbuffer.so
	go build -a .

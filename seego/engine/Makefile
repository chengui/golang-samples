.PHONY: build run all

all: build run

run:
	LD_LIBRARY_PATH=`pwd`/lib ./cgo-engine

build:
	gcc -fpic -shared -I./include ./include/*.cc -o lib/libengine.so
	go build -a .

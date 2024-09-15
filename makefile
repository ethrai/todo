.PHONY: clean build run all test

build:
	go build -o ./build/todo ./main.go

run: build
	./build/todo


.DEFAULT_GOAL := build

build:
	go build -o bin/websshket main.go

clean:
	rm bin/*

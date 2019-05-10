binname := scriba

.PHONY: run
run:build
	sudo ./$(binname)

.PHONY: test
test:
	go test

.PHONY: build
build: test
	go build .

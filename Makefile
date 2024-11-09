.PHONY: build
build:
	CGO_ENABLED=0 go build -o ./build/main

.PHONY: lint
lint:
	golangci-lint run -v

.PHONY: tests
tests:
	go test -v -race \
		-covermode atomic \
		-coverprofile coverage.out \
		./... -json > report.json
	go tool cover -func coverage.out

.PHONY: run
run:
	./build/main

.PHONY: clean
clean:
	rm -rf build/
	rm coverage.out report.json

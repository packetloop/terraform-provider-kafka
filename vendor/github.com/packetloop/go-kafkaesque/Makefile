all: run coverage test

.PHONY: test
test:
	# Ignore doc_test.go because example would make a HTTP request
	# and fail.
	go test -race -v ./... -tags doc

.PHONY: coverage
coverage: test
	go get github.com/axw/gocov/gocov
	gocov test -tags doc | gocov report

.PHONY: run
run:
	go get -v ./...
	go fmt ./...
	go vet ./...
	dep ensure

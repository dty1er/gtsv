.PHONY: test
test:
	go test -v

.PHONY: cov
cov:
	go test -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html
	open cover.html


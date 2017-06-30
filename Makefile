.PHONY: compile
compile:
	export CGO_ENABLED=0
	export GOOS=linux
	go build -a -installsuffix cgo -o growser .

.PHONY: test
test:
	go test -v .

.PHONY: build
build:
	docker-compose build

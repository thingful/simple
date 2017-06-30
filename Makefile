.PHONY: compile
compile:
	export CGO_ENABLED=0
	export GOOS=linux
	go build -a -installsuffix cgo -o growser .

.PHONY: build
build:
	docker build -t thingful/growser .

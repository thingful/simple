.PHONY: build
build:
	export CGO_ENABLED=0
	export GOOS=linux
	go build -a -installsuffix cgo -o growser

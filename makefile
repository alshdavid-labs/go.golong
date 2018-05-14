GOPATH=${CURDIR}
export GOPATH=${CURDIR}

all: install build
build:
	go install -v ./src/app
install:
	go env GOPATH
	go get -d -v ./...
watch:
	go get github.com/canthefason/go-watcher
	go install github.com/canthefason/go-watcher/cmd/watcher
	./bin/watcher -watch ./app -run ./src/app 

# For windows
# choco install make gnuwin32-coreutils.install -y
# You may need to add C:\Program Files (x86)\GnuWin32\bin to your path

all: get build

.PHONY: build
build: clean compile

.PHONY: env
env:
	go env

.PHONY: clean
clean:
	rm -r ./bin || true

.PHONY: compile
compile:
	mkdir bin
	cp ./src/.env ./bin/.env
	cd src && env GOBIN=${CURDIR}/bin go install ./cmd/servid

.PHONY: get
get:
	cd src && go get -d -v ./...

.PHONY: run
run:
	cd src && go run ./cmd/servid/main.go

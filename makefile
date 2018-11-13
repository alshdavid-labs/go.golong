default: build

PHONY: build
build:
	mkdir -p bin
	cd bin && go build ../cmd/httpd/main.go

PHONY: run
run:
	go run ./cmd/httpd/main.go

PHONY: watch
watch:
	gin -b ./bin/gin-bin run ./cmd/httpd/main.go
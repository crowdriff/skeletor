version=0.1.0
name=skeletor
vcs=github.com/crowdriff/$(name)

.PHONY: all

# To cross compile for linux on mac, build go-linux cross compiler first using
# cd /usr/local/go/src
# sudo GOOS=linux GOARCH=amd64 CGO_ENABLED=0 ./make.bash --no-clean

all:
	@echo "make <cmd>"
	@echo ""
	@echo "commands:"
	@echo "  build         - build the dist binary"
	@echo "  build_linux   - force build a linux binary"
	@echo "  clean         - clean the dist build"
	@echo "  coverage      - generates test coverage report"
	@echo "  deps          - download dependencies"
	@echo "  install       - run go install"
	@echo "  run           - run with live reload"
	@echo "  test          - standard go test"
	@echo "  tools         - go gets a bunch of tools for dev"
	@echo "  update_deps   - update dependency list"

build: clean
	@go vet ./...
	@golint ./...
	@go build -o ./bin/$(name)-$(version).bin main.go

build_linux: clean
	@go vet ./...
	@golint ./...
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/$(name)-$(version).bin main.go

clean:
	@rm -rf ./bin

coverage:
	go test -p=1 -cover -v ./...

deps:
	@glock sync -n $(vcs) < Glockfile

install:
	@go install ./...

run:
	@fresh

test:
	@ginkgo -r

tools:
	go get github.com/robfig/glock
	go get github.com/pilu/fresh
	go get github.com/golang/lint/golint
	go get github.com/onsi/ginkgo
	go get github.com/onsi/gomega

update_deps:
	@glock save -n $(vcs) > Glockfile

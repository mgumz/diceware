VERSION=$(shell cat VERSION)
BUILD_DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GIT_HASH=$(shell git rev-parse HEAD)

BUILDS=linux.amd64 linux.386 linux.arm64 linux.mips64 windows.amd64.exe freebsd.amd64 darwin.amd64 darwin.arm64
BINARIES=$(addprefix bin/diceware-$(VERSION)., $(BUILDS))

LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.BuildDate=$(BUILD_DATE) -X main.GitHash=$(GIT_HASH)"

bin/diceware: bin
	go build -v -o $@ .

releases: $(BINARIES)

compile-analysis:
	go build -gcflags '-m' .

code-quality:
	-go vet ./
	-gofmt -s -d ./
	-golint ./
	-gocyclo ./
	-ineffassign ./

bin/diceware-$(VERSION).linux.%: bin
	env GOOS=linux GOARCH=$* CGO_ENABLED=0 go build $(LDFLAGS) -o $@ .

bin/diceware-$(VERSION).darwin.%: bin
	env GOOS=darwin GOARCH=$* CGO_ENABLED=0 go build $(LDFLAGS) -o $@ .

bin/diceware-$(VERSION).windows.%.exe: bin
	env GOOS=windows GOARCH=$* CGO_ENABLED=0 go build $(LDFLAGS) -o $@ .

bin/diceware-$(VERSION).freebsd.%: bin
	env GOOS=freebsd GOARCH=$* CGO_ENABLED=0 go build $(LDFLAGS) -o $@ .

bin:
	mkdir $@

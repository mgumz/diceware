VERSION=0.1.0
BUILD_DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GIT_HASH=$(shell git rev-parse HEAD)

RELEASES=bin/diceware-$(VERSION).linux.amd64 \
		 bin/diceware-$(VERSION).linux.arm64 \
		 bin/diceware-$(VERSION).linux.mips64 \
		 bin/diceware-$(VERSION).windows.amd64.exe \
		 bin/diceware-$(VERSION).darwin.amd64


LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.BuildDate=$(BUILD_DATE) -X main.GitHash=$(GIT_HASH)"


releases: $(RELEASES)

bin/diceware-$(VERSION).linux.mips64: bin
	env GOOS=linux GOARCH=mips64 CGO_ENABLED=0 go build $(LDFLAGS) -o $@

bin/diceware-$(VERSION).linux.amd64: bin
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o $@

bin/diceware-$(VERSION).linux.arm64: bin
	env GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build $(LDFLAGS) -o $@

bin/diceware-$(VERSION).windows.amd64.exe: bin
	env GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o $@

bin/diceware-$(VERSION).darwin.amd64: bin
	env GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o $@

bin:
	mkdir $@

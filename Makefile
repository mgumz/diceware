VERSION=$(shell cat VERSION)
BUILD_DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GIT_HASH=$(shell git rev-parse HEAD)

BUILDS=linux.amd64 linux.386 linux.arm64 linux.mips64 windows.amd64.exe freebsd.amd64 darwin.amd64 darwin.arm64
BINARIES=$(addprefix bin/diceware-$(VERSION)., $(BUILDS))

LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.BuildDate=$(BUILD_DATE) -X main.GitHash=$(GIT_HASH)"

bin/diceware: bin .
	go build -v -o $@ .

releases: $(BINARIES)

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


######################################################
## dev related

compile-analysis: cmd/mtr-exporter
	go build -gcflags '-m' ./$^

fmt:
	-gofmt -s -d ./

code-quality: report-cyclo report-mispell report-lint report-ineffassign report-staticcheck

report-cyclo:
	@echo '####################################################################'
	gocyclo .
report-mispell:
	@echo '####################################################################'
	misspell ./...
report-lint:
	@echo '####################################################################'
	golint ./...
report-ineffassign:
	@echo '####################################################################'
	ineffassign ./...
report-vet:
	@echo '####################################################################'
	go vet ./...
report-staticcheck:
	@echo '####################################################################'
	staticcheck ./...

fetch-report-tools:
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	go install github.com/client9/misspell/cmd/misspell@latest
	go install github.com/gordonklaus/ineffassign@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install golang.org/x/lint/golint@latest

.PHONY: bin/diceware
uname_p := $(shell uname -p) # store the output of the command in a variable

build: pre_build
	mkdir -p bin
	go build -o ./bin/lscdl ./cmd/lscdl

pre_build:
	./set_version.sh
	go mod tidy
	mkdir -p ./bin

build_dlv:
	go get github.com/go-delve/delve/cmd/dlv@latest
	mkdir -p ~/go/bin/dlv
	go build -o ~/bin/dlv.$(uname_p) github.com/go-delve/delve/cmd/dlv
	ln -sf ~/bin/dlv.$(uname_p) ~/bin/dlv

# Use the following on m1:
# alias make='/usr/bin/arch -arch arm64 /usr/bin/make'
debug:
	go build -gcflags "all=-N -l" -o ./bin/lscdl.debug ./cmd/lscdl
	~/bin/dlv --headless --listen=:2345 --api-version=2 --accept-multiclient exec ./bin/lscdl.debug

#debug_test:
#	~/go/bin/dlv --headless --listen=:2345 --api-version=2 --accept-multiclient test ./pkg/lscdl

fmt:
	gofmt -w .
	goimports -w .
	gci write .

test: sec lint

sec:
	gosec ./...

lint:
	golangci-lint run

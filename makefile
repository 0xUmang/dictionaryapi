GOPATH ?= $(shell go env GOPATH)
GOBIN ?= $(GOPATH)/bin

build-maclinux:
	export PATH="$(shell go env GOPATH/bin):$$PATH"
	export API_KEY=b48e3347-81bd-4b38-9da5-2d72399817fd
	go build -o $(GOBIN)/dictionaryapi .

build-windows:
	go build -o dictionaryapi.exe .

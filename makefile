GOBIN ?=$(GOPATH)/bin

include .env
export

build-maclinux:
	export API_KEY=b48e3347-81bd-4b38-9da5-2d72399817fd
	go build -o $(GOBIN)/dictionaryapi .

build-windows:
	go build -o dictionaryapi.exe .
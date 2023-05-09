.DEFAULT_GOAL := build

define template
package main\n\nfunc main() {}
endef

name = $(subst -,_,$(problem_name))
gen:
	echo "create problem folder $(name)"
	mkdir -p ./$(name)
	echo "$(template)" > ./$(name)/$(name).go
	touch ./$(name)/$(name)_test.go

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build hello.go
.PHONY:build

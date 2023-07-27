.DEFAULT_GOAL := build

define template
package main

func main() {}

func example() {
	// TODO
}
endef

define test_template
package main

import (
	"reflect"
	"testing"
)

func Test_example(t *testing.T) {
	tests := []struct {
		// nums []int
		// i    []int
	}{
		// {[]int{8, 1, 2, 2, 3}, []int{4, 0, 1, 1, 3}},
	}
	for _, test := range tests {
		result := example(/*test.nums*/)
		if !reflect.DeepEqual(result, test.i) {
			t.Errorf("Total was incorrect, got: %v, want: %v.", result, test.i)
		}
	}
}
endef


export template
export test_template

name = $(subst -,_,$(problem_name))
gen:
	echo "create problem folder $(name)"
	mkdir -p ./$(name)
	@echo "$$template" > ./$(name)/$(name).go
	@echo "$$test_template" > ./$(name)/$(name)_test.go

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

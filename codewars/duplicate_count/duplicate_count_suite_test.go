package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDuplicateCount(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DuplicateCount Suite")
}

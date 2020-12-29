package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Split Strings", func() {
	It("Basic tests", func() {
		Expect(Solution("abc")).To(Equal([]string{"ab", "c_"}))
		Expect(Solution("dawsd")).To(Equal([]string{"da", "ws", "d_"}))
		Expect(Solution("awsaws")).To(Equal([]string{"aw", "sa", "ws"}))
	})
})

func TestSplitStrings(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SplitStrings Suite")
}

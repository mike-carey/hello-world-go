package hello_world_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	foo "github.com/mike-carey/hello-world-golang"
)

var _ = Describe("HelloWorld", func() {
	It("should return foo", func() {
		Expect(foo.Foo()).Should(Equal("foo"))
	})
})

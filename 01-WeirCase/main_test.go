package main

import (
	. "github.com/onsi/ginkgo"
	"testing"
)
import . "github.com/onsi/gomega"

var _ = Describe("Sample Test Cases:", func() {
	It("Should return the correct values", func() {
		Expect(toWeirdCase("abc def")).To(Equal("AbC DeF"))
		Expect(toWeirdCase("ABC")).To(Equal("AbC"))
		Expect(toWeirdCase("This is a test Looks like you passed")).To(Equal("ThIs Is A TeSt LoOkS LiKe YoU PaSsEd"))
	})
})


func Test_toWeirdCase(t *testing.T) {

}

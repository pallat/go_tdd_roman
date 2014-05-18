package roman_test

import (
	. "tdd/roman"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Roman", func() {
  It("should return I when input is 1", func() {
	var roman = Roman{Number:1}
	Expect(roman.GetRoman()).To(Equal("I"))
  })
  It("should return II when input is 2", func() {
	var roman = Roman{Number:2}
	Expect(roman.GetRoman()).To(Equal("II"))
  })
  It("should return III when input is 3", func() {
	var roman = Roman{Number:3}
	Expect(roman.GetRoman()).To(Equal("III"))
  })
  It("should return IV when input is 4", func() {
	var roman = Roman{Number:4}
	Expect(roman.GetRoman()).To(Equal("IV"))
  })
  It("should return V when input is 5", func() {
	var roman = Roman{Number:5}
	Expect(roman.GetRoman()).To(Equal("V"))
  })
  It("should return VI when input is 6", func() {
	var roman = Roman{Number:6}
	Expect(roman.GetRoman()).To(Equal("VI"))
  })
  It("should return VII when input is 7", func() {
	var roman = Roman{Number:7}
	Expect(roman.GetRoman()).To(Equal("VII"))
  })
  It("should return VIII when input is 8", func() {
	var roman = Roman{Number:8}
	Expect(roman.GetRoman()).To(Equal("VIII"))
  })
  It("should return IX when input is 9", func() {
	var roman = Roman{Number:9}
	Expect(roman.GetRoman()).To(Equal("IX"))
  })
  It("should return X when input is 10", func() {
	var roman = Roman{Number:10}
	Expect(roman.GetRoman()).To(Equal("X"))
  })
})

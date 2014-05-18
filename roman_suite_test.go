package roman_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRoman(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Roman Suite")
}

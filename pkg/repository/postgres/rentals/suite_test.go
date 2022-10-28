package rentals_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRentals(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rentals Suite")
}

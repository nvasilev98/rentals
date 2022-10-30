package rentals_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	dbClient *sql.DB
	mock     sqlmock.Sqlmock
)

var _ = BeforeSuite(func() {
	var err error
	dbClient, mock, err = sqlmock.New()
	Expect(err).ToNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	mock.ExpectClose()
	Expect(dbClient.Close()).To(Succeed())
})

func TestRentals(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rentals Suite")
}

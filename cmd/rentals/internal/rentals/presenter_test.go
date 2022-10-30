package rentals_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/nvasilev98/rentals/cmd/rentals/internal/rentals"
	"github.com/nvasilev98/rentals/cmd/rentals/internal/rentals/mocks"
	"github.com/nvasilev98/rentals/pkg/api"
	r "github.com/nvasilev98/rentals/pkg/repository/postgres/rentals"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Presenter", func() {
	var (
		gomockCtrl     *gomock.Controller
		mockRentalRepo *mocks.MockRentalRepository
		presenter      *rentals.Presenter
		recorder       *httptest.ResponseRecorder
		mockContext    *gin.Context
	)

	BeforeEach(func() {
		gomockCtrl, _ = gomock.WithContext(context.Background(), GinkgoT())
		mockRentalRepo = mocks.NewMockRentalRepository(gomockCtrl)
		presenter = rentals.NewPresenter(mockRentalRepo)
		recorder = httptest.NewRecorder()
		mockContext, _ = gin.CreateTestContext(recorder)
	})

	AfterEach(func() {
		gomockCtrl.Finish()
	})

	When("id parameter is missing", func() {
		BeforeEach(func() {
			mockContext.Request, _ = http.NewRequest(http.MethodGet, gomock.Any().String(), nil)
		})

		It("should return http.StatusBadRequest code", func() {
			presenter.RetrieveRentalByID(mockContext)
			Expect(mockContext.Writer.Status()).To(Equal(http.StatusBadRequest))
			errResp := api.ErrorResponse{}
			Expect(json.Unmarshal(recorder.Body.Bytes(), &errResp)).To(Succeed())
			Expect(errResp.Error.Message).To(Equal("missing id parameter"))
		})
	})

	When("retrieving rental by id from repository fails", func() {
		BeforeEach(func() {
			mockContext.Request, _ = http.NewRequest(http.MethodGet, gomock.Any().String(), nil)
			mockContext.Params = []gin.Param{{Key: "id", Value: "1"}}
			mockRentalRepo.EXPECT().RetrieveRentalByID(gomock.Any(), gomock.Any()).Return(r.Model{}, errors.New("err"))
		})

		It("should return http.StatusInternalServerError code", func() {
			presenter.RetrieveRentalByID(mockContext)
			Expect(mockContext.Writer.Status()).To(Equal(http.StatusInternalServerError))
			errResp := api.ErrorResponse{}
			Expect(json.Unmarshal(recorder.Body.Bytes(), &errResp)).To(Succeed())
			Expect(errResp.Error.Message).To(Equal("failed to retrieve rental by id"))
		})
	})

	When("retrieving rental by id succeeds", func() {
		const (
			id   = 1
			name = "test"
		)

		BeforeEach(func() {
			mockContext.Request, _ = http.NewRequest(http.MethodGet, gomock.Any().String(), nil)
			mockContext.Params = []gin.Param{{Key: "id", Value: "1"}}
			mockRentalRepo.EXPECT().RetrieveRentalByID(gomock.Any(), gomock.Any()).Return(r.Model{ID: id, Name: name}, nil)
		})

		It("should return http.StatusOK code", func() {
			presenter.RetrieveRentalByID(mockContext)
			Expect(mockContext.Writer.Status()).To(Equal(http.StatusOK))
			rentalResp := rentals.RentalResponse{}
			Expect(json.Unmarshal(recorder.Body.Bytes(), &rentalResp)).To(Succeed())
			Expect(rentalResp.ID).To(Equal(id))
			Expect(rentalResp.Name).To(Equal(name))
		})
	})

	When("retrieving rentals from repository fails", func() {
		BeforeEach(func() {
			mockContext.Request, _ = http.NewRequest(http.MethodGet, gomock.Any().String(), nil)
			mockRentalRepo.EXPECT().RetrieveRentals(gomock.Any(), gomock.Any()).Return([]r.Model{}, errors.New("err"))
		})

		It("should return http.StatusInternalServerError code", func() {
			presenter.RetrieveRentals(mockContext)
			Expect(mockContext.Writer.Status()).To(Equal(http.StatusInternalServerError))
			errResp := api.ErrorResponse{}
			Expect(json.Unmarshal(recorder.Body.Bytes(), &errResp)).To(Succeed())
			Expect(errResp.Error.Message).To(Equal("failed to retrieve rentals"))
		})
	})

	When("retrieving rentals succeeds", func() {
		const (
			id   = 1
			name = "test"
		)

		BeforeEach(func() {
			mockContext.Request, _ = http.NewRequest(http.MethodGet, gomock.Any().String(), nil)
			mockRentalRepo.EXPECT().RetrieveRentals(gomock.Any(), gomock.Any()).Return([]r.Model{{ID: id, Name: name}}, nil)
		})

		It("should return http.StatusOK code", func() {
			presenter.RetrieveRentals(mockContext)
			Expect(mockContext.Writer.Status()).To(Equal(http.StatusOK))
			rentalResp := rentals.RentalsResponse{}
			Expect(json.Unmarshal(recorder.Body.Bytes(), &rentalResp)).To(Succeed())
			Expect(rentalResp.Rentals[0].ID).To(Equal(id))
			Expect(rentalResp.Rentals[0].Name).To(Equal(name))
		})
	})
})

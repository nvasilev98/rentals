package rentals_test

import (
	"context"
	"errors"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nvasilev98/rentals/pkg/repository/postgres/rentals"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const expectedSelectRentals = `SELECT 
							r.id, name, description, type, vehicle_make, vehicle_model, vehicle_year,
							vehicle_length, sleeps, primary_image_url, price_per_day, home_city, home_state,
							home_zip, home_country, lat, lng, user_id, first_name, last_name
							FROM rentals r
							LEFT JOIN users u
							ON r.user_id = u.id`

var _ = Describe("Rentals", func() {
	AfterEach(func() {
		Expect(mock.ExpectationsWereMet()).To(Succeed())
	})

	Context("NewRepository", func() {
		When("preparing select rental by id statement fails", func() {
			BeforeEach(func() {
				mock.ExpectPrepare(expectedSelectRentals).WillReturnError(errors.New("err"))
			})

			It("should return an error", func() {
				_, err := rentals.NewRepository(dbClient)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Context("RetrieveRentalByID", func() {
		var (
			prepare    *sqlmock.ExpectedPrepare
			repository *rentals.Repository
			err        error
			ctx        context.Context
		)

		BeforeEach(func() {
			prepare = mock.ExpectPrepare(expectedSelectRentals)
			repository, err = rentals.NewRepository(dbClient)
			Expect(err).ToNot(HaveOccurred())
			ctx = context.Background()
		})

		AfterEach(func() {
			Expect(repository.Close()).To(Succeed())
		})

		When("executing prepared query statement fails", func() {
			BeforeEach(func() {
				prepare.ExpectQuery().WithArgs(sqlmock.AnyArg()).WillReturnError(errors.New("err"))
			})

			It("should return an error", func() {
				_, err := repository.RetrieveRentalByID(ctx, "id")
				Expect(err).To(HaveOccurred())
			})
		})

		When("retrieving a rental by a given id", func() {
			testFields := []string{"r.id", "name", "description", "type", "vehicle_make", "vehicle_model", "vehicle_year",
				"vehicle_length", "sleeps", "primary_image_url", "price_per_day", "home_city", "home_state",
				"home_zip", "home_country", "lat", "lng", "user_id", "first_name", "last_name"}
			expectedRental := rentals.Model{
				ID: 1, Name: "name", Description: "description", Type: "type", VehicleMake: "maker",
				VehicleModel: "model", VehicleYear: 2, VehicleLength: 123.3, Sleeps: 3, PrimaryImageURL: "URL",
				PricePerDay: 10, HomeCity: "city", HomeState: "state", HomeZIP: "ZIP", HomeCountry: "country",
				LAT: 123.2, LNG: 456.1, UserID: 3, FirstName: "first-name", LastName: "last-name"}

			BeforeEach(func() {
				mockRows := mock.NewRows(testFields).
					AddRow(expectedRental.ID, expectedRental.Name, expectedRental.Description,
						expectedRental.Type, expectedRental.VehicleMake, expectedRental.VehicleModel,
						expectedRental.VehicleYear, expectedRental.VehicleLength, expectedRental.Sleeps,
						expectedRental.PrimaryImageURL, expectedRental.PricePerDay, expectedRental.HomeCity,
						expectedRental.HomeState, expectedRental.HomeZIP, expectedRental.HomeCountry,
						expectedRental.LAT, expectedRental.LNG, expectedRental.UserID,
						expectedRental.FirstName, expectedRental.LastName)
				prepare.ExpectQuery().WithArgs(sqlmock.AnyArg()).WillReturnRows(mockRows)
			})

			It("should succeeds", func() {
				rental, err := repository.RetrieveRentalByID(ctx, "1")
				Expect(err).ToNot(HaveOccurred())
				Expect(rental).To(Equal(expectedRental))
			})
		})
	})
})

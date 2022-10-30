package rentals

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nvasilev98/rentals/pkg/api"
	"github.com/nvasilev98/rentals/pkg/repository/postgres/rentals"
)

//go:generate mockgen --source=presenter.go --destination mocks/presenter.go --package mocks

type RentalRepository interface {
	RetrieveRentalByID(ctx context.Context, id string) (rentals.Model, error)
	RetrieveRentals(ctx context.Context, query map[string][]string) ([]rentals.Model, error)
}

type Presenter struct {
	rentalRepository RentalRepository
}

// NewPresenter is a constructor function
func NewPresenter(rentalRepository RentalRepository) *Presenter {
	return &Presenter{
		rentalRepository: rentalRepository,
	}
}

// RetrieveRentalByID retrieves a rental by a given id
func (p *Presenter) RetrieveRentalByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, api.NewErrorResponse("missing id parameter"))
		return
	}

	rental, err := p.rentalRepository.RetrieveRentalByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, api.NewErrorResponse("failed to retrieve rental by id"))
		return
	}

	ctx.JSON(http.StatusOK, toRentalResponse(rental))
}

func (p *Presenter) RetrieveRentals(ctx *gin.Context) {
	queryParams := ctx.Request.URL.Query()
	rentals, err := p.rentalRepository.RetrieveRentals(ctx, queryParams)
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusInternalServerError, api.NewErrorResponse("failed to retrieve rentals"))
		return
	}

	ctx.JSON(http.StatusOK, rentals)
}

func toRentalResponse(rental rentals.Model) RentalResponse {
	price := PriceResponse{Day: rental.PricePerDay}
	location := LocationResponse{
		HomeCity:    rental.HomeCity,
		HomeState:   rental.HomeState,
		HomeZIP:     rental.HomeZIP,
		HomeCountry: rental.HomeCountry,
		LAT:         rental.LAT,
		LNG:         rental.LNG,
	}
	user := UserResponse{
		ID:        rental.ID,
		FirstName: rental.FirstName,
		LastName:  rental.LastName,
	}

	return RentalResponse{
		ID:              rental.ID,
		Name:            rental.Name,
		Description:     rental.Description,
		Type:            rental.Type,
		VehicleMake:     rental.VehicleMake,
		VehicleModel:    rental.VehicleModel,
		VehicleYear:     rental.VehicleYear,
		VehicleLength:   rental.VehicleLength,
		Sleeps:          rental.Sleeps,
		PrimaryImageURL: rental.PrimaryImageURL,
		Price:           price,
		Location:        location,
		User:            user,
	}
}

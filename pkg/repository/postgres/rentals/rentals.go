package rentals

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository struct {
	db                   *sql.DB
	selectRentalByIDStmt *sql.Stmt
}

func NewRepository(db *sql.DB) (*Repository, error) {
	selectRentalByIDStmt, err := db.Prepare(selectRentalByID)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare select rental by id statement: %w", err)
	}

	return &Repository{
		db:                   db,
		selectRentalByIDStmt: selectRentalByIDStmt,
	}, nil
}

func (r *Repository) RetrieveRentalByID(ctx context.Context, id string) (Model, error) {
	row := r.selectRentalByIDStmt.QueryRowContext(ctx, id)
	var rental Model
	err := row.Scan(
		&rental.ID,
		&rental.Name,
		&rental.Description,
		&rental.Type,
		&rental.VehicleMake,
		&rental.VehicleModel,
		&rental.VehicleYear,
		&rental.VehicleLength,
		&rental.Sleeps,
		&rental.PrimaryImageURL,
		&rental.PricePerDay,
		&rental.HomeCity,
		&rental.HomeState,
		&rental.HomeZIP,
		&rental.HomeCountry,
		&rental.LAT,
		&rental.LNG,
		&rental.UserID,
		&rental.FirstName,
		&rental.LastName,
	)
	if err != nil {
		return Model{}, fmt.Errorf("failed to scan row: %w", err)
	}

	return rental, nil
}

func (r *Repository) Close() error {
	if err := r.selectRentalByIDStmt.Close(); err != nil {
		return fmt.Errorf("failed to close select rental by id statement: %w", err)
	}

	return nil
}

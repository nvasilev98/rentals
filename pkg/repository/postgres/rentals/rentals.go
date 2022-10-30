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

// NewRepository is a constructor function
func NewRepository(db *sql.DB) (*Repository, error) {
	selectRentalByIDStmt, err := db.Prepare(fmt.Sprintf("%s WHERE r.id=$1", selectRentals))
	if err != nil {
		return nil, fmt.Errorf("failed to prepare select rental by id statement: %w", err)
	}

	return &Repository{
		db:                   db,
		selectRentalByIDStmt: selectRentalByIDStmt,
	}, nil
}

// RetrieveRentalByID retrieves rental by a given id from repository
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

// RetrieveRentals retrieves filtered rentals by adding clauses to the query
func (r *Repository) RetrieveRentals(ctx context.Context, clauses map[string][]string) ([]Model, error) {
	rows, err := r.db.QueryContext(ctx, buildSQLQuery(selectRentals, clauses))
	if err != nil {
		return nil, fmt.Errorf("failed to execute select rentals query: %w", err)
	}
	defer rows.Close()

	rentals := make([]Model, 0)
	for rows.Next() {
		var rental Model
		err := rows.Scan(
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
			return nil, fmt.Errorf("failed to scan a row: %w", err)
		}

		rentals = append(rentals, rental)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("failed while iterating over rows: %w", rows.Err())
	}

	return rentals, nil
}

// Close closes statements for repository
func (r *Repository) Close() error {
	if err := r.selectRentalByIDStmt.Close(); err != nil {
		return fmt.Errorf("failed to close select rental by id statement: %w", err)
	}

	return nil
}

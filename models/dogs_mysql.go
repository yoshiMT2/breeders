package models

import (
	"context"
	"time"
)

func (d *DogBreed) AllDogBreeds() ([]*DogBreed, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()

	query := `select * from dog_breeds order by breed`

	var breeds []*DogBreed

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var b DogBreed
		err := rows.Scan(
			&b.ID,
			&b.Breed,
			&b.WeightLowLbs,
			&b.WeightHighLbs,
			&b.Lifespan,
			&b.Details,
			&b.AlternateNames,
			&b.GeographicOrigin,
		)
		if err != nil {
			return nil, err
		}
		breeds = append(breeds, &b)
	}
	return breeds, nil
}

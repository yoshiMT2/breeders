package models

import (
	"database/sql"
	"time"
)

var db *sql.DB

type Models struct {
	DogBreed DogBreed
}

func New(conn *sql.DB) *Models {
	db = conn

	return &Models{
		DogBreed: DogBreed{},
	}
}

type DogBreed struct {
	ID               int    `json:"id"`
	Breed            string `json:"breed"`
	WeightLowLbs     int    `json:"weigth_low_lbs"`
	WeightHighLbs    int    `json:"weigth_high_lbs"`
	Lifespan         int    `json:"averate_lifespan"`
	Details          string `json:"details"`
	AlternateNames   string `json:"alternate_names"`
	GeographicOrigin string `json:"geographic_origin"`
}

func (d *DogBreed) All() ([]*DogBreed, error) {
	return d.AllDogBreeds()
}

type CatBreed struct {
	ID               int    `json:"id"`
	Breed            string `json:"breed"`
	WeightLowLbs     int    `json:"weigth_low_lbs"`
	WeightHighLbs    int    `json:"weigth_high_lbs"`
	Lifespan         int    `json:"averate_lifespan"`
	Details          string `json:"details"`
	AlternateNames   string `json:"alternate_names"`
	GeographicOrigin string `json:"geographic_origin"`
}

type Dog struct {
	ID               int       `json:"id"`
	DogName          string    `json:"dog_name"`
	BreedID          int       `json:"breed_id"`
	BreederID        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	SpyaedOrNeutered int       `json:"spayed_neutered"`
	Description      string    `json:"desicription"`
	Weight           int       `json:"weight"`
	Breed            DogBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}

type Cat struct {
	ID               int       `json:"id"`
	CatName          string    `json:"cat_name"`
	BreedID          int       `json:"breed_id"`
	BreederID        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	SpayedOrNeutered int       `json:"spayed_neutered"`
	Description      string    `json:"desicription"`
	Weight           int       `json:"weight"`
	Breed            CatBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}

type Breeder struct {
	ID          int         `json:"id"`
	BreederName string      `json:"breder_name"`
	Address     string      `json:"address"`
	City        string      `json:"city"`
	ProvState   string      `json:"prov_state"`
	Country     string      `json:"country"`
	Zip         string      `json:"zip"`
	Phone       string      `json:"phone"`
	Email       string      `json:"email"`
	Active      int         `json:"active"`
	DogBreeds   []*DogBreed `json:"dog_breeds"`
	CatBreeds   []*CatBreed `json:"cat_breeds"`
}

type Pet struct {
	Species      string `json:"species"`
	Breed        string `json:"breed"`
	MinWeight    int    `json:"min_weight"`
	MaxWeight    int    `json:"max_weight"`
	Desicription string `json:"description"`
	LifeSpan     int    `json:"lifespan"`
}

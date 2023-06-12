package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type TripRepository interface {
	FindTrips() ([]models.Trip, error)
	GetTrip(Id int) (models.Trip, error)
	GetUpdateId(Id int) (models.Trip, error)
	GetCountryId(Id int) (models.CountriesResponse, error)
	DeleteTrip(trip models.Trip, ID int) (models.Trip, error)
	CreateTrip(trip models.Trip) (models.Trip, error)
	UpdateTrip(trip models.Trip) (models.Trip, error)
}

func RepositoryTrip(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTrips() ([]models.Trip, error) {
	var trips []models.Trip
	err := r.db.Preload("Country").Find(&trips).Error

	return trips, err
}

func (r *repository) GetTrip(Id int) (models.Trip, error) {
	var trip models.Trip

	err := r.db.First(&trip, Id).Error

	return trip, err
}

func (r *repository) GetUpdateId(Id int) (models.Trip, error) {
	var Countries models.Trip
	err := r.db.Preload("Country").First(&Countries, Id).Error
	
	return Countries, err
}

func (r *repository) GetCountryId(Id int) (models.CountriesResponse, error) {
	var Countries models.CountriesResponse
	err := r.db.First(&Countries, Id).Error

	return Countries, err
}

func (r *repository) DeleteTrip(trip models.Trip, ID int) (models.Trip, error) {
	err := r.db.Delete(&trip, ID).Scan(&trip).Error

	return trip, err
}

func (r *repository) CreateTrip(trip models.Trip) (models.Trip, error) {
	err := r.db.Create(&trip).Error

	return trip, err
}

func (r *repository) UpdateTrip(trip models.Trip) (models.Trip, error) {
	err := r.db.Save(&trip).Error

	return trip, err
}

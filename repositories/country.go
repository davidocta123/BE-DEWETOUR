package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type CountryRepository interface {
	FindCountrys() ([]models.Country, error)
	GetCountry(Id int) (models.Country, error)
	CreateCountry(country models.Country) (models.Country, error)
	UpdateCountry(country models.Country) (models.Country, error)
	DeleteCountry(country models.Country, ID int) (models.Country, error)
}

func RepositoryCountry(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCountrys() ([]models.Country, error) {
	var countrys []models.Country
	err := r.db.Find(&countrys).Error

	return countrys, err
}

func (r *repository) GetCountry(Id int) (models.Country, error) {
	var country models.Country

	err := r.db.First(&country, Id).Error

	return country, err
}

func (r *repository) CreateCountry(country models.Country) (models.Country, error) {
	err := r.db.Create(&country).Error

	return country, err
}

func (r *repository) UpdateCountry(country models.Country) (models.Country, error) {
	err := r.db.Save(&country).Error

	return country, err
}

func (r *repository) DeleteCountry(country models.Country, ID int) (models.Country, error) {
	err := r.db.Delete(&country, ID).Scan(&country).Error

	return country, err
}

package kamar

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Kamar, error)
	FindByNamaKamar(nama string) ([]Kamar, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Kamar, error) {
	var kamar []Kamar
	err := r.db.Find(&kamar).Error

	return kamar, err
}

func (r *repository) FindByNamaKamar(nama_kamar string) ([]Kamar, error) {
	var kamar []Kamar
	// err := r.db.Where("title = ?", title).First(&users).Error

	value := fmt.Sprintf("%%%s%%", nama_kamar)
	err := r.db.Where("nama LIKE ?", value).Find(&kamar).Error

	// err := r.db.Where("email_user LIKE ?", name_product).Find(&allProducts).Error

	return kamar, err
}

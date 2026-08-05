package repository

import (
	"gorm.io/gorm"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/model"
)

type DriverRepository struct {
	db *gorm.DB
}

func NewDriverRepository(
	db *gorm.DB,
) *DriverRepository{
	return &DriverRepository{
		db: db,
	}
}

func (r *DriverRepository) Create(
	driver *model.Driver,
) error {
	return r.db.Create(driver).Error
}

func (r *DriverRepository) GetByID(
	id uint,
) (*model.Driver, error){
	var driver model.Driver
	if err := r.db.First(&driver, id).Error; err != nil {
		return nil, err
	}
	return &driver, nil
}

func (r *DriverRepository) Update(
	driver *model.Driver,
) error {
	return r.db.Save(driver).Error
}	
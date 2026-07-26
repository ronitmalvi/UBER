package repository

import ("github.com/ronitmalvi/UBER/ride-service-go/internal/model"
		"gorm.io/gorm"
)
type RideRepository interface {
	Create(ride *model.Ride) error
	GetByID(id uint) (*model.Ride, error)
	Update(ride *model.Ride) error
	Delete(id uint) error
}

type rideRepository struct {

	db *gorm.DB
}

func NewRideRepository(db *gorm.DB) RideRepository {
	return &rideRepository{db: db}
}

func (r *rideRepository) Create(
	ride *model.Ride,
) error {

	return r.db.Create(ride).Error
}

func (r *rideRepository) GetByID(
	id uint,
) (*model.Ride, error) {

	var ride model.Ride

	err := r.db.First(
		&ride,
		id,
	).Error

	return &ride, err
}

func (r *rideRepository) Update(
	ride *model.Ride,
) error {

	return r.db.Save(ride).Error
}

func (r *rideRepository) Delete(
	id uint,
) error {

	return r.db.Delete(
		&model.Ride{},
		id,
	).Error
}
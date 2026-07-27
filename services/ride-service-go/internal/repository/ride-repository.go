//Repositories answer one question:
//How do I store or retrieve data?

package repository

import ("github.com/ronitmalvi/UBER/ride-service-go/internal/model"
		"gorm.io/gorm"
)
type RideRepository struct {                       // A RideRepository contains a database connection.
    db *gorm.DB
}

func NewRideRepository(
    db *gorm.DB,
) *RideRepository {

    return &RideRepository{
        db: db,
    }
}

func (r *RideRepository) Create(                        //r *RideRepository -> This function belongs to RideRepository.
    ride *model.Ride,
) error {

    return r.db.Create(ride).Error
}
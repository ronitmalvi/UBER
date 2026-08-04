//Repositories answer one question:
//How do I store or retrieve data?

package repository

import ("github.com/ronitmalvi/UBER/ride-service-go/internal/model"
		"gorm.io/gorm"
        "fmt"
)
type RideRepository struct {                       // A RideRepository contains a database connection.
    db *gorm.DB
}

func NewRideRepository(                             //Constructor function for RideRepository.
    db *gorm.DB,
) *RideRepository {

    return &RideRepository{
        db: db,
    }
}

func (r *RideRepository) Create(                        //r *RideRepository -> This function belongs to RideRepository.
    ride *model.Ride,
) error {
    fmt.Println("Creating ride in the database...(repository layer)")
    return r.db.Create(ride).Error
}

func (r *RideRepository) GetByID(
    id uint,
) (*model.Ride, error){
    var ride model.Ride
    if err := r.db.First(&ride, id).Error; err != nil {
        return nil, err
    }
    return &ride, nil
}
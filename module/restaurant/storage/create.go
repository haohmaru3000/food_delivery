package restaurantstorage

import (
	"context"

	restaurantmodel "github.com/0xThomas3000/food_delivery/module/restaurant/model"
)

func (s *sqlStore) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

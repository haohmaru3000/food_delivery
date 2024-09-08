package rstlikestorage

import (
	"context"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/modules/restaurantlike/model"
)

func (s *sqlStore) Create(ctx context.Context, data *rstlikemodel.Like) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	// db.Exec("Update restaurants SET liked_count = liked_count + 1 where id = ?", data.RestaurantId) -> Very bad way to do

	return nil
}

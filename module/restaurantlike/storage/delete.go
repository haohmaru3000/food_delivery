package rstlikestorage

import (
	"context"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/module/restaurantlike/model"
)

func (s *sqlStore) Delete(ctx context.Context, data *rstlikemodel.LikeDelete) error {
	db := s.db.Table(rstlikemodel.LikeDelete{}.TableName()).
		Where("user_id = ? and restaurant_id = ?", data.UserId, data.RestaurantId).
		Delete(&data)

	if err := db.Error; err != nil {
		return common.ErrDB(err)
	} else if db.RowsAffected < 1 {
		return rstlikemodel.ErrCannotDislikeRestaurant(err)
	}

	return nil
}

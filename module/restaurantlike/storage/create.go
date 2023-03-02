package restaurantlikestorage

import (
	"context"
	"github.com/0xThomas3000/food_delivery/common"
	restaurantlikemodel "github.com/0xThomas3000/food_delivery/module/restaurantlike/model"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantlikemodel.Like) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

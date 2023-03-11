package rstlikestorage

import (
	"context"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/module/restaurantlike/model"
)

func (s *sqlStore) Create(ctx context.Context, data *rstlikemodel.Like) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

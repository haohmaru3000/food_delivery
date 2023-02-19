package restaurantstorage

import (
	"context"
	"gorm.io/gorm"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/module/restaurant/model"
)

func (s *sqlStore) FindDataWithCondition(context context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant

	// First: phát sinh 2 TH
	// 1: nếu ko tìm thấy -> return lỗi gorm.ErrRecordNotFound
	// 2: lỗi khác (DB parse sai key)
	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}

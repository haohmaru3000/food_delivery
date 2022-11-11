package biz

import (
	"context"
	"github.com/0xThomas3000/food_delivery/common"
	restaurantmodel "github.com/0xThomas3000/food_delivery/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	Delete(context context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(context context.Context, id int) error {
	oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})
	// Nếu để kiểm tra email có valid ko?
	// 1. Check lỗi có phải gorm.ErrRecordNotFound (nếu đúng thì ok để làm tiếp [chưa ai sd this email])
	// 2. Care only lỗi này: chưa chắc ok (ở t.đ find data xuống, chưa chắc db đang work ok, như too many connection...)
	if err != nil {
		return common.ErrEntityNotFound(restaurantmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil) // nil vì ko có error gốc
	}

	if err := biz.store.Delete(context, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, nil) // TH có lỗi gốc
	}
	return nil
}

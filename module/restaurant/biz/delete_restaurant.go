package restaurantbiz

import (
	"context"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/module/restaurant/model"
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
	// thêm vào field dưới -> để đảm bảo requester chỉ xoá dc nhà hàng của chính họ(họ phải sở hữu nhà hàng đó)
	requester common.Requester // là Interface -> có thể mock nó để unit test rất dễ dàng
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore, requester common.Requester) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{
		store:     store,
		requester: requester,
	}
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

	if oldData.UserId != biz.requester.GetUserId() {
		return common.ErrNoPermission(nil)
	}

	if err := biz.store.Delete(context, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, nil) // TH có lỗi gốc
	}
	return nil
}

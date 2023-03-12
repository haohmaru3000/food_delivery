package rstlikebiz

import (
	"context"

	"github.com/0xThomas3000/food_delivery/module/restaurantlike/model"
)

// Need 2 one more Store:
// +	Checking if this Restaurant exists -> If yes, the user can like it?
// +  Checking whether this Restaurant has been liked or not? (-1:31:0)
type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *rstlikemodel.Like) error
}

type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store: store,
	}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context, data *rstlikemodel.Like) error {
	err := biz.store.Create(ctx, data)

	if err != nil {
		return rstlikemodel.ErrCannotLikeRestaurant(err)
	}

	return nil
}

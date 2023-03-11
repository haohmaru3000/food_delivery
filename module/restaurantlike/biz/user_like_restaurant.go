package rstlikebiz

import (
	"context"

	"github.com/0xThomas3000/food_delivery/module/restaurantlike/model"
)

// Need one more Store (for Finding a Restaurant which has been liked yet?) (-1:31:0)
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

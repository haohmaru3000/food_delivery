package rstlikebiz

import (
	"context"

	"github.com/0xThomas3000/food_delivery/module/restaurantlike/model"
)

type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, data *rstlikemodel.LikeDelete) error
}

type userDislikeRestaurantBiz struct {
	store UserDislikeRestaurantStore
}

func NewUserDislikeRestaurantBiz(store UserDislikeRestaurantStore) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{
		store: store,
	}
}

func (biz *userDislikeRestaurantBiz) DislikeRestaurant(ctx context.Context, data *rstlikemodel.LikeDelete) error {
	err := biz.store.Delete(ctx, data)

	if err != nil {
		return rstlikemodel.ErrCannotDislikeRestaurant(err)
	}

	return nil
}

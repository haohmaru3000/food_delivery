package rstlikebiz

import (
	"context"
	"log"

	"github.com/0xThomas3000/food_delivery/components/asyncjob"
	"github.com/0xThomas3000/food_delivery/module/restaurantlike/model"
)

type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, data *rstlikemodel.LikeDelete) error
}

type DecLikedCountResStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userDislikeRestaurantBiz struct {
	store    UserDislikeRestaurantStore
	decStore DecLikedCountResStore
}

func NewUserDislikeRestaurantBiz(store UserDislikeRestaurantStore, decStore DecLikedCountResStore) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{
		store:    store,
		decStore: decStore,
	}
}

func (biz *userDislikeRestaurantBiz) DislikeRestaurant(ctx context.Context, data *rstlikemodel.LikeDelete) error {
	err := biz.store.Delete(ctx, data)

	if err != nil {
		return rstlikemodel.ErrCannotDislikeRestaurant(err)
	}

	// Side effect
	j := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.decStore.DecreaseLikeCount(ctx, data.RestaurantId)
	})

	if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
		log.Println(err)
	}

	return nil
}

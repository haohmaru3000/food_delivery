package rstlikebiz

import (
	"context"
	"log"

	"github.com/0xThomas3000/food_delivery/common"
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

	go func() {
		defer common.AppRecover() // To ensure the service won't get broken if we have a Crash/Panic
		// time.Sleep(time.Second * 3) // To demonstrate the API below won't get blocked as we put 'side effect flow' in Goroutine
		if err := biz.decStore.DecreaseLikeCount(ctx, data.RestaurantId); err != nil {
			log.Println(err)
		}
	}()

	return nil
}

package rstlikebiz

import (
	"context"
	"log"
	"time"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/module/restaurantlike/model"
)

// Need 2 one more Store:
// +	Checking if this Restaurant exists -> If yes, the user can like it?
// +  Checking whether this Restaurant has been liked or not? (-1:31:0)
type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *rstlikemodel.Like) error
}

type IncLikedCountResStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store    UserLikeRestaurantStore
	incStore IncLikedCountResStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore, incStore IncLikedCountResStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store:    store,
		incStore: incStore,
	}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context, data *rstlikemodel.Like) error {
	err := biz.store.Create(ctx, data)

	if err != nil {
		return rstlikemodel.ErrCannotLikeRestaurant(err)
	}

	go func() {
		defer common.AppRecover() // To ensure the service won't get broken if we have a Crash/Panic
		// time.Sleep(time.Second * 3) // To demonstrate the API below won't get blocked as we put 'side effect flow' in Goroutine
		if err := biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId); err != nil {
			log.Println(err)

			// Retry to restart the service again three times if it's broken
			for i := 0; i < 3; i++ {
				err := biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
				if err == nil {
					break
				}
				time.Sleep(time.Second * 3) // Time sleep between every Retry
			}
		}
	}()

	return nil
}

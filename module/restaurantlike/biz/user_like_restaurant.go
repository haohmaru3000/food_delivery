package rstlikebiz

import (
	"context"
	"log"

	"github.com/0xThomas3000/food_delivery/components/asyncjob"
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

	// Side effect
	j := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
	})

	if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
		log.Println(err)
	}

	return nil
}

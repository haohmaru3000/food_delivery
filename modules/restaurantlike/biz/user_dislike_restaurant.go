package rstlikebiz

import (
	"context"
	"log"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/modules/restaurantlike/model"
	"github.com/0xThomas3000/food_delivery/pubsub"
)

type UserDislikeRestaurantStore interface {
	Delete(ctx context.Context, data *rstlikemodel.LikeDelete) error
}

// type DecLikedCountResStore interface {
// 	DecreaseLikeCount(ctx context.Context, id int) error
// }

type userDislikeRestaurantBiz struct {
	store UserDislikeRestaurantStore
	// decStore DecLikedCountResStore
	ps pubsub.Pubsub
}

func NewUserDislikeRestaurantBiz(
	store UserDislikeRestaurantStore,
	// decStore DecLikedCountResStore,
	ps pubsub.Pubsub,
) *userDislikeRestaurantBiz {
	return &userDislikeRestaurantBiz{
		store: store,
		// decStore: decStore,
		ps: ps,
	}
}

func (biz *userDislikeRestaurantBiz) DislikeRestaurant(ctx context.Context, data *rstlikemodel.LikeDelete) error {
	err := biz.store.Delete(ctx, data)

	if err != nil {
		return rstlikemodel.ErrCannotDislikeRestaurant(err)
	}

	// Send message
	if err := biz.ps.Publish(ctx, common.TopicUserDislikeRestaurant, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}

	// Side effect
	// j := asyncjob.NewJob(func(ctx context.Context) error {
	// 	return biz.decStore.DecreaseLikeCount(ctx, data.RestaurantId)
	// })

	// if err := asyncjob.NewGroup(true, j).Run(ctx); err != nil {
	// 	log.Println(err)
	// }

	return nil
}

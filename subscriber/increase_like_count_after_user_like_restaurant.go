package subscriber

import (
	"context"
	"log"

	"github.com/0xThomas3000/food_delivery/components/appctx"
	"github.com/0xThomas3000/food_delivery/modules/restaurant/storage"
	"github.com/0xThomas3000/food_delivery/pubsub"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	// GetUserId() int
}

// func IncreaseLikecountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
// 	c, _ := appCtx.GetPubsub().Subscribe(ctx, common.TopicUserLikeRestaurant)

// 	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())

// 	go func() {
// 		defer common.AppRecover()
// 		for {
// 			msg := <-c
// 			likeData := msg.Data().(HasRestaurantId)
// 			_ = store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
// 		}
// 	}()
// }

func IncreaseLikecountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like-count after user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)

			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}

func PushNotificationWhenUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Push notification after user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasRestaurantId)
			log.Println("Push notification when user likes restaurant_id: ", likeData.GetRestaurantId())

			// Ex: Send message (if token failed or sth...). [Sect12: 01:07:00]

			return nil
		},
	}
}

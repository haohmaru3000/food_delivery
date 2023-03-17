package subscriber

import (
	"context"
	"log"

	"github.com/0xThomas3000/food_delivery/components/appctx"
	"github.com/0xThomas3000/food_delivery/module/restaurant/storage"
	"github.com/0xThomas3000/food_delivery/pubsub"
)

// func DecreaseLikecountAfterUserDislikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
// 	c, _ := appCtx.GetPubsub().Subscribe(ctx, common.TopicUserDislikeRestaurant)

// 	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())

// 	go func() {
// 		defer common.AppRecover()
// 		for {
// 			msg := <-c
// 			dislikeData := msg.Data().(HasRestaurantId)
// 			_ = store.DecreaseLikeCount(ctx, dislikeData.GetRestaurantId())
// 		}
// 	}()
// }

func DecreaseLikecountAfterUserDislikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease like-count after user dislikes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
			dislikeData := message.Data().(HasRestaurantId)
			return store.DecreaseLikeCount(ctx, dislikeData.GetRestaurantId())
		},
	}
}

func PushNotificationWhenUserDislikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Push notification after user dislikes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			dislikeData := message.Data().(HasRestaurantId)
			log.Println("Push notification when user dislikes restaurant_id: ", dislikeData.GetRestaurantId())

			return nil
		},
	}
}

/* The structure code for 'engine.go' file */
// Collect and Reduce all codes to create Subscriber -> into this code for New dev
// func RunSomething(ac appctx.AppContext) func(ctx context.Context, msg *pubsub.Message) error {
// 	return func(ctx context.Context, msg *pubsub.Message) error {
// 		store := restaurantstorage.NewSQLStore(ac.GetMainDBConnection())
// 		dislikeData := msg.Data().(HasRestaurantId)
// 		return store.DecreaseLikeCount(ctx, dislikeData.GetRestaurantId())
// 	}
// }

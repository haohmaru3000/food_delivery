package subscriber

import (
	"context"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/components/appctx"
	"github.com/0xThomas3000/food_delivery/module/restaurant/storage"
)

func DecreaseLikecountAfterUserDislikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, common.TopicUserDislikeRestaurant)

	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())

	go func() {
		defer common.AppRecover()
		for {
			msg := <-c
			dislikeData := msg.Data().(HasRestaurantId)
			_ = store.DecreaseLikeCount(ctx, dislikeData.GetRestaurantId())
		}
	}()
}

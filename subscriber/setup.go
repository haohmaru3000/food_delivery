package subscriber

import (
	"context"

	"github.com/0xThomas3000/food_delivery/components/appctx"
)

func Setup(appCtx appctx.AppContext, ctx context.Context) {
	IncreaseLikecountAfterUserLikeRestaurant(appCtx, ctx)
	DecreaseLikecountAfterUserDislikeRestaurant(appCtx, ctx)
}

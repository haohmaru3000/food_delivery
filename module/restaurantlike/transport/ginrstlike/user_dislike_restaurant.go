package ginrstlike

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/components/appctx"
	"github.com/0xThomas3000/food_delivery/module/restaurantlike/biz"
	"github.com/0xThomas3000/food_delivery/module/restaurantlike/model"
	"github.com/0xThomas3000/food_delivery/module/restaurantlike/storage"
)

// DELETE /v1/restaurants/:id/dislike
func UserDislikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := rstlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		// decStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := rstlikebiz.NewUserDislikeRestaurantBiz(store, appCtx.GetPubsub())

		likeDelete := &rstlikemodel.LikeDelete{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		if err := biz.DislikeRestaurant(c.Request.Context(), likeDelete); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

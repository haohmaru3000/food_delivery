package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/component/appctx"
	restaurantbiz "github.com/0xThomas3000/food_delivery/module/restaurant/biz"
	"github.com/0xThomas3000/food_delivery/module/restaurant/model"
	"github.com/0xThomas3000/food_delivery/module/restaurant/storage"
	"github.com/0xThomas3000/food_delivery/module/restaurantlike/storage"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter.Status = []int{1} // Set quyền

		store := restaurantstorage.NewSQLStore(db)
		likeStore := restaurantlikestorage.NewSQLStore(db)
		biz := restaurantbiz.NewListRestaurantBiz(store, likeStore)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false) // Chưa login => false
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}

package ginrestaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/components/appctx"
	"github.com/0xThomas3000/food_delivery/module/restaurant/biz"
	"github.com/0xThomas3000/food_delivery/module/restaurant/model"
	"github.com/0xThomas3000/food_delivery/module/restaurant/repository"
	"github.com/0xThomas3000/food_delivery/module/restaurant/storage"
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
		// likeStore := rstlikestorage.NewSQLStore(db)
		repo := restaurantrepo.NewListRestaurantRepo(store)
		biz := restaurantbiz.NewListRestaurantBiz(repo)

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

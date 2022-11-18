package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/component/appctx"
	restaurantbiz "github.com/0xThomas3000/food_delivery/modules/restaurant/biz"
	"github.com/0xThomas3000/food_delivery/modules/restaurant/model"
	"github.com/0xThomas3000/food_delivery/modules/restaurant/storage"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		// Crash error (panic xảy ra trong 1 Goroutine): needs to be treated as "normal error"
		// go func() {
		// 	defer common.AppRecover()

		// 	arr := []int{}
		// 	log.Println(arr[0])

		// }()

		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}

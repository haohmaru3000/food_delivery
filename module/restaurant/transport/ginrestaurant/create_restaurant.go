package ginrestaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/components/appctx"
	"github.com/0xThomas3000/food_delivery/module/restaurant/biz"
	"github.com/0xThomas3000/food_delivery/module/restaurant/model"
	"github.com/0xThomas3000/food_delivery/module/restaurant/storage"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		requester := c.MustGet(common.CurrentUser).(common.Requester) // Lấy ra người gửi truy vấn này lên

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

		data.UserId = requester.GetUserId()

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}

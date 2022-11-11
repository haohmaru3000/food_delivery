package ginrestaurant

import (
	"log"
	"net/http"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/component/appctx"
	restaurantbiz "github.com/0xThomas3000/food_delivery/module/restaurant/biz"
	restaurantmodel "github.com/0xThomas3000/food_delivery/module/restaurant/model"
	restaurantstorage "github.com/0xThomas3000/food_delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		// Crash error: needs to be treated as "normal error"
		arr := []int{}
		log.Println(arr[0])

		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err)) // xử lý lỗi
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			// Lỗi bên business -> chỉ cần quăng lỗi về luôn ko cần xử lý bọc lại (vì đã biết chắc return lỗi của mình)
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}

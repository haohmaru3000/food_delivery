package ginrestaurant

import (
	"net/http"
	"strconv"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/component/appctx"
	restaurantbiz "github.com/0xThomas3000/food_delivery/module/restaurant/biz"
	restaurantstorage "github.com/0xThomas3000/food_delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err)) // ko panic(err) vì đây là lỗi gốc
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			panic(err) // Chắc chắn là tầng business đã xử lý rồi
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}

package ginrestaurant

import (
	"fmt"
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
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered:", r)
				}
			}()

			arr := []int{}
			log.Println(arr[0])

		}()

		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}

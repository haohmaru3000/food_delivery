package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/0xThomas3000/food_delivery/components/appctx"
	"github.com/0xThomas3000/food_delivery/middleware"
	restaurantmodel "github.com/0xThomas3000/food_delivery/module/restaurant/model"
	"github.com/0xThomas3000/food_delivery/module/restaurant/transport/ginrestaurant"
	"github.com/0xThomas3000/food_delivery/module/upload/uploadtransport/ginupload"
	"github.com/0xThomas3000/food_delivery/module/user/transport/ginuser"
)

func SetupRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	v1.POST("/upload", ginupload.Upload(appContext))

	v1.POST("/register", ginuser.Register(appContext))
	v1.POST("/authenticate", ginuser.Login(appContext))
	// Hàm Profile() sẽ không thể lấy dc Context liên quan tới User nếu không thông qua "Middleware" dc ĐN trước
	v1.GET("/profile", middleware.RequiredAuth(appContext), ginuser.Profile(appContext))

	// ROUTER GROUP for restaurants
	restaurants := v1.Group("/restaurants", middleware.RequiredAuth(appContext))
	{
		restaurants.POST("/", ginrestaurant.CreateRestaurant(appContext))
	}

	restaurants.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var data restaurantmodel.Restaurant
		appContext.GetMainDBConnection().Where("id = ?", id).First(&data)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	restaurants.GET("/", ginrestaurant.ListRestaurant(appContext))

	restaurants.PATCH("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var data restaurantmodel.RestaurantUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		appContext.GetMainDBConnection().Where("id = ?", id).Updates(&data)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))
}

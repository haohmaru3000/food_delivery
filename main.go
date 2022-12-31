package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/0xThomas3000/food_delivery/component/appctx"
	"github.com/0xThomas3000/food_delivery/component/uploadprovider"
	"github.com/0xThomas3000/food_delivery/middleware"
	"github.com/0xThomas3000/food_delivery/modules/restaurant/model"
	"github.com/0xThomas3000/food_delivery/modules/restaurant/transport/ginrestaurant"
	"github.com/0xThomas3000/food_delivery/modules/upload/uploadtransport/ginupload"
	"github.com/0xThomas3000/food_delivery/modules/user/transport/ginuser"
	"github.com/0xThomas3000/food_delivery/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUserName, config.DBUserPassword, config.DBHost, config.DBPort, config.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(config.S3BucketName, config.S3Region, config.S3APIKey, config.S3SecretKey, config.S3Domain)

	appContext := appctx.NewAppContext(db, s3Provider)

	r := gin.Default()
	r.Use(middleware.Recover(appContext))

	// Đăng ký link cho cái static để hiển thị hình
	r.Static("/static", "./static") // Đi search mục "static" => gin sẽ kiếm thư mục "static" để đọc

	v1 := r.Group("/v1")
	v1.POST("/upload", ginupload.Upload(appContext))

	v1.POST("/register", ginuser.Register(appContext))

	// ROUTER GROUP for restaurants
	restaurants := v1.Group("/restaurants")
	restaurants.POST("/", ginrestaurant.CreateRestaurant(appContext))

	restaurants.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var data restaurantmodel.Restaurant
		db.Where("id = ?", id).First(&data)
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

		db.Where("id = ?", id).Updates(&data)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))

	r.Run()
}

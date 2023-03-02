package main

import (
	"github.com/gin-gonic/gin"

	"github.com/0xThomas3000/food_delivery/component/appctx"
	"github.com/0xThomas3000/food_delivery/middleware"
	"github.com/0xThomas3000/food_delivery/module/user/transport/ginuser"
)

func setupAdminRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	admin := v1.Group("/admin",
		middleware.RequiredAuth(appContext),
		middleware.RoleRequired(appContext, "admin", "mod"),
	)

	{
		admin.GET("/profile", ginuser.Profile(appContext))
	}
}

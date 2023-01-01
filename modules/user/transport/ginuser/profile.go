package ginuser

import (
	"net/http"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/component/appctx"
	"github.com/gin-gonic/gin"
)

func Profile(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(user))
	}
}

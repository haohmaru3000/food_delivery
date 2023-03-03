package middleware

import (
	"errors"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/components/appctx"
	"github.com/gin-gonic/gin"
)

func RoleRequired(appCtx appctx.AppContext, allowRoles ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		hasFound := false

		for _, item := range allowRoles {
			// Look up the current user's role in the DB, compare it to the "allowRoles" params
			if u.GetRole() == item {
				hasFound = true
				break
			}
		}

		if !hasFound {
			panic(common.ErrNoPermission(errors.New("invalid role user")))
		}

		c.Next()
	}
}

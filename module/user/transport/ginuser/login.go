package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/0xThomas3000/food_delivery/common"
	"github.com/0xThomas3000/food_delivery/components/appctx"
	"github.com/0xThomas3000/food_delivery/components/hasher"
	"github.com/0xThomas3000/food_delivery/components/tokenprovider/jwt"
	userbiz "github.com/0xThomas3000/food_delivery/module/user/biz"
	usermodel "github.com/0xThomas3000/food_delivery/module/user/model"
	userstore "github.com/0xThomas3000/food_delivery/module/user/store"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		business := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}

package appctx

import (
	"gorm.io/gorm"

	"github.com/0xThomas3000/food_delivery/component/uploadprovider"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	db             *gorm.DB
	uploadprovider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, uploadprovider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{
		db:             db,
		uploadprovider: uploadprovider,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadprovider
}

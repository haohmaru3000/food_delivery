package appctx

import (
	"gorm.io/gorm"

	"github.com/0xThomas3000/food_delivery/component/uploadprovider"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
}

// A struct that implements "AppContext" interface
type appCtx struct {
	db             *gorm.DB
	uploadprovider uploadprovider.UploadProvider
}

// Setters for setting 2 fields of "db", "uploadprovider"
func NewAppContext(db *gorm.DB, uploadprovider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{
		db:             db,
		uploadprovider: uploadprovider,
	}
}

// Getters to get 2 fields of "db", "uploadprovider"
func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadprovider
}

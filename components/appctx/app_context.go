package appctx

import (
	"gorm.io/gorm"

	"github.com/0xThomas3000/food_delivery/components/uploadprovider"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
}

// A struct that implements "AppContext" interface
type appCtx struct {
	db             *gorm.DB
	uploadprovider uploadprovider.UploadProvider
	secretKey      string
}

// Setters for setting 2 fields of "db", "uploadprovider"
func NewAppContext(db *gorm.DB, uploadprovider uploadprovider.UploadProvider, secretKey string) *appCtx {
	return &appCtx{
		db:             db,
		uploadprovider: uploadprovider,
		secretKey:      secretKey,
	}
}

// Getters to get 2 fields of "db", "uploadprovider"
func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadprovider
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}

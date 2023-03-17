package appctx

import (
	"gorm.io/gorm"

	"github.com/0xThomas3000/food_delivery/components/uploadprovider"
	"github.com/0xThomas3000/food_delivery/pubsub"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubsub() pubsub.Pubsub
}

// A struct that implements "AppContext" interface
type appCtx struct {
	db             *gorm.DB
	uploadprovider uploadprovider.UploadProvider
	secretKey      string
	ps             pubsub.Pubsub
}

// Setters for setting 2 fields of "db", "uploadprovider"
func NewAppContext(
	db *gorm.DB, uploadprovider uploadprovider.UploadProvider,
	secretKey string, ps pubsub.Pubsub,
) *appCtx {
	return &appCtx{
		db:             db,
		uploadprovider: uploadprovider,
		secretKey:      secretKey,
		ps:             ps,
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

func (ctx *appCtx) GetPubsub() pubsub.Pubsub {
	return ctx.ps
}

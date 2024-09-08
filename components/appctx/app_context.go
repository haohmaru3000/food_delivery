package appctx

import (
	"gorm.io/gorm"

	"github.com/0xThomas3000/food_delivery/components/uploadprovider"
	"github.com/0xThomas3000/food_delivery/pubsub"
	"github.com/0xThomas3000/food_delivery/skio"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubsub() pubsub.Pubsub
	GetRealtimeEngine() skio.RealtimeEngine
}

// A struct that implements "AppContext" interface
type appCtx struct {
	db             *gorm.DB
	uploadprovider uploadprovider.UploadProvider
	secretKey      string
	ps             pubsub.Pubsub
	rtEngine       skio.RealtimeEngine
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
func (ctx *appCtx) GetMainDBConnection() *gorm.DB                 { return ctx.db }
func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider { return ctx.uploadprovider }
func (ctx *appCtx) SecretKey() string                             { return ctx.secretKey }
func (ctx *appCtx) GetPubsub() pubsub.Pubsub                      { return ctx.ps }
func (ctx *appCtx) GetRealtimeEngine() skio.RealtimeEngine        { return ctx.rtEngine }

func (ctx *appCtx) SetRealtimeEngine(rt skio.RealtimeEngine) { ctx.rtEngine = rt }

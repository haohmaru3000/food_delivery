package skuser

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
	"gorm.io/gorm"

	"github.com/0xThomas3000/food_delivery/common"
)

type SmallAppContext interface {
	GetMainDBConnection() *gorm.DB
}

type LocationData struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func OnUserUpdateLocation(
	appCtx SmallAppContext,
	requester common.Requester,
) func(s socketio.Conn, location LocationData) {
	return func(s socketio.Conn, location LocationData) {
		// In case you want to store to DB?
		// appCtx.GetMainDBConnection()
		// ...

		// location belong to user ???
		log.Println("User update location: user id is", requester.GetUserId(), "at location", location)
	}
}

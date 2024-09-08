package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/0xThomas3000/food_delivery/components/appctx"
	"github.com/0xThomas3000/food_delivery/components/uploadprovider"
	"github.com/0xThomas3000/food_delivery/middleware"
	"github.com/0xThomas3000/food_delivery/pubsub/localpb"
	"github.com/0xThomas3000/food_delivery/routes"
	"github.com/0xThomas3000/food_delivery/skio"
	"github.com/0xThomas3000/food_delivery/subscriber"
	"github.com/0xThomas3000/food_delivery/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUserName, config.DBUserPassword, config.DBHost, config.DBPort, config.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(
		config.S3BucketName,
		config.S3Region,
		config.S3APIKey,
		config.S3SecretKey,
		config.S3Domain,
	)

	secretKey := config.SecretKey
	ps := localpb.NewPubSub()
	appContext := appctx.NewAppContext(db, s3Provider, secretKey, ps)

	// Setup subscribers
	// subscriber.Setup(appContext, context.Background())
	_ = subscriber.NewEngine(appContext).Start()

	r := gin.Default()

	// Read 'demo.html' on Server when Client enters 'localhost:8080/demo' in URL
	r.StaticFS("/demo", http.Dir("./client"))

	r.Use(middleware.Recover(appContext))

	// Đăng ký link cho cái static để hiển thị hình
	r.Static("/static", "./static") // Đi search mục "static" => gin sẽ kiếm thư mục "static" để đọc

	// CRUD
	v1 := r.Group("/v1")

	routes.SetupRoute(appContext, v1)
	routes.SetupAdminRoute(appContext, v1)

	// startSocketIOServer(r, appContext)

	rtEngine := skio.NewEngine()
	appContext.SetRealtimeEngine(rtEngine)

	_ = rtEngine.Run(appContext, r)

	/* Will fail if placing here because of EOF err */
	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}

// func startSocketIOServer(engine *gin.Engine, appCtx appctx.AppContext) {
// 	server := socketio.NewServer(&engineio.Options{
// 		// Two types of Transport: 1. websocket, 2. long polling
// 		Transports: []transport.Transport{websocket.Default},
// 	})

// 	// Listen to a namespace "/"(default), then trigger calback function:
// 	// 	+ Everytime a Client gets connected to Server, then log that connection.
// 	server.OnConnect("/", func(serverSocket socketio.Conn) error {
// 		// s.SetContext("")
// 		fmt.Println("Socket connected:", serverSocket.ID(), " IP:", serverSocket.RemoteAddr())

// 		// Whatever sockets come in, join them into 'Shipper' room
// 		// At Client side, they won't know that they were pushed into 'Shipper'
// 		serverSocket.Join("Shipper")

// 		//server.BroadcastToRoom("/", "Shipper", "test", "Hello G05")

// 		return nil
// 	})

// 	// go func() {
// 	// 	// Use for range in Channel. Until Channel is closed, the loop will break
// 	// 	for range time.NewTicker(time.Second).C {
// 	// 		server.BroadcastToRoom("/", "Shipper", "test", "Hi room !")
// 	// 	}
// 	// }()

// 	// This way to trigger a callback when error, not return error via Payload or Header...
// 	// ex: error when parsing-data process, error at library or error at Transport layer
// 	server.OnError("/", func(serverSocket socketio.Conn, e error) {
// 		fmt.Println("Meet error:", e)
// 	})

// 	// Client is disconnected to Server
// 	server.OnDisconnect("/", func(serverSocket socketio.Conn, reason string) {
// 		fmt.Println("Closed", reason)
// 		// Remove socket from socket engine (from app context)
// 	})

// 	server.OnEvent("/", "authenticate", func(serverSocket socketio.Conn, token string) {

// 		// Validate token
// 		// If false: s.Close(), and return

// 		// If true
// 		// => UserId
// 		// Fetch db find user by Id
// 		// Here: s belongs to who? (user_id)
// 		// We need a map[user_id][]socketio.Conn

// 		db := appCtx.GetMainDBConnection()
// 		store := userstorage.NewSQLStore(db)
// 		//
// 		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
// 		//
// 		payload, err := tokenProvider.Validate(token)

// 		if err != nil {
// 			serverSocket.Emit("authentication_failed", err.Error())
// 			serverSocket.Close()
// 			return
// 		}
// 		//
// 		user, err := store.FindUser(context.Background(), map[string]interface{}{"id": payload.UserId})
// 		//
// 		if err != nil {
// 			serverSocket.Emit("authentication_failed", err.Error())
// 			serverSocket.Close()
// 			return
// 		}

// 		if user.Status == 0 {
// 			serverSocket.Emit("authentication_failed", errors.New("you has been banned/deleted"))
// 			serverSocket.Close()
// 			return
// 		}

// 		user.Mask(false)

// 		serverSocket.Emit("your_profile", user)
// 	})

// 	// Server listening to Client, then event "test" (like Topic in pubsub)
// 	server.OnEvent("/", "test", func(serverSocket socketio.Conn, msg interface{}) {
// 		log.Println("test:", msg)
// 	})

// 	type Person struct {
// 		Name string `json:"name"`
// 		Age  int    `json:"age"`
// 	}

// 	// Golang's reflex will automatically detect p
// 	// If p is of type Struct, then it'll unmarshall byte buffer(is Json string) -> to Person struct
// 	server.OnEvent("/", "notice", func(serverSocket socketio.Conn, p Person) {
// 		fmt.Println("Server receives notice:", p.Name, p.Age)

// 		p.Age = 40
// 		serverSocket.Emit("notice", p)

// 	})

// 	//server.OnEvent("/chat", "msg", func(serverSocket socketio.Conn, msg string) string {
// 	//	s.SetContext(msg)
// 	//	return "recv " + msg
// 	//})
// 	//
// 	//server.OnEvent("/", "bye", func(serverSocket socketio.Conn) string {
// 	//	last := s.Context().(string)
// 	//	s.Emit("bye", last)
// 	//	s.Close()
// 	//	return last
// 	//})
// 	//
// 	//server.OnEvent("/", "noteSumit", func(serverSocket socketio.Conn) string {
// 	//	last := s.Context().(string)
// 	//	s.Emit("bye", last)
// 	//	s.Close()
// 	//	return last
// 	//})

// 	// go server.Serve() // Listen to web-socket
// 	go func() {
// 		if err := server.Serve(); err != nil {
// 			log.Fatalf("SocketIO listen error: %s\n", err)
// 		}
// 	}()
// 	defer server.Close()

// 	// Ko thể đi thẳng lên ở web-socket
// 	// => Có thể mượn đường http server để upgrade lên, vì web-socket based on http
// 	// ex: gọi sth vào path, then trigger HTTP handler
// 	engine.GET("/socket.io/*any", gin.WrapH(server))
// 	engine.POST("/socket.io/*any", gin.WrapH(server))
// 	engine.StaticFS("/demo", http.Dir("./client"))

// 	if err := engine.Run(":8080"); err != nil {
// 		log.Fatal("Failed run app: ", err)
// 	}
// }

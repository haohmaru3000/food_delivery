package main

import (
	"context"
	"log"
	"time"

	"github.com/0xThomas3000/food_delivery/pubsub"
	"github.com/0xThomas3000/food_delivery/pubsub/localpb"
)

func main() {
	var localPS pubsub.Pubsub = localpb.NewPubSub()

	//chn := pubsub.Topic("OrderCreated")
	var topic pubsub.Topic = "OrderCreated"

	sub1, close1 := localPS.Subscribe(context.Background(), topic)
	sub2, _ := localPS.Subscribe(context.Background(), topic)

	localPS.Publish(context.Background(), topic, pubsub.NewMessage(1))
	localPS.Publish(context.Background(), topic, pubsub.NewMessage(2))

	go func() {
		for {
			log.Println("Con1:", (<-sub1).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	go func() {
		for {
			log.Println("Con2:", (<-sub2).Data())
			time.Sleep(time.Millisecond * 400)
		}
	}()

	time.Sleep(time.Second * 3)
	close1()
	//close2()

	localPS.Publish(context.Background(), topic, pubsub.NewMessage(3))

	time.Sleep(time.Second * 2)
}

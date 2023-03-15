package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/0xThomas3000/food_delivery/components/asyncjob"
)

func main() {
	job1 := asyncjob.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second)
		log.Println("I am job 1")

		// return nil
		return errors.New("something went wrong at job 1")
	})

	job1.SetRetryDurations([]time.Duration{time.Second * 3}) // We want job1 to only run Retry once

	if err := job1.Execute(context.Background()); err != nil {
		log.Println(job1.State(), err)

		// Retry
		for {
			if err := job1.Retry(context.Background()); err != nil {
				log.Println(err)
			}
			if job1.State() == asyncjob.StateRetryFailed || job1.State() == asyncjob.StateCompleted {
				log.Println(job1.State())
				break
			}
		}
	}
}

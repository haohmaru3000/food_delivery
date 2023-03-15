package asyncjob

import (
	"context"
	"time"
)

// Job Requirement:
// 1. Job can do something (handler)
// 2. Job can retry
// 	2.1 Config retry times and duration
// 3. Should be stateful (if job is still running, we need its state to know more about it)
// 4. We should have job manager to manage jobs (*)

type Job interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	State() JobState
	SetRetryDurations(times []time.Duration)
}

const (
	defaultMaxTimeout = time.Second * 10
)

var (
	// Khi job failed -> lần retry đầu hẹn 4s, lần 2 hẹn thêm 2s(4s + 2s = 6s), lần 3 = 7s
	defaultRetryTime = []time.Duration{time.Second, time.Second * 2, time.Second * 4}
)

type JobHandler func(ctx context.Context) error

type JobState int

const (
	StateInit JobState = iota // Enum in Golang (iota = 0)
	StateRunning
	StateFailed
	StateTimeout
	StateCompleted
	StateRetryFailed
)

// Replace enum above with number by []string below
func (js JobState) String() string {
	return []string{"Init", "Running", "Failed", "Timeout", "Completed", "RetryFailed"}[js]
}

type jobConfig struct {
	MaxTimeout time.Duration
	Retries    []time.Duration
}

type job struct {
	config     jobConfig
	handler    JobHandler
	state      JobState
	retryIndex int
	stopChan   chan bool
}

func NewJob(handler JobHandler) *job {
	j := job{
		config: jobConfig{
			MaxTimeout: defaultMaxTimeout,
			Retries:    defaultRetryTime,
		},
		handler:    handler,
		retryIndex: -1, // Not yet run Retry
		state:      StateInit,
		stopChan:   make(chan bool),
	}

	return &j
}

func (j *job) Execute(ctx context.Context) error {
	j.state = StateRunning // Assign State to "Running"

	var err error = j.handler(ctx)

	if err != nil {
		j.state = StateFailed
		return err
	}

	j.state = StateCompleted

	return nil

	/** To cancel or interfere into this Flow */
	//ch := make(chan error)
	//ctxJob, doneFunc := context.WithCancel(ctx)

	//go func() {
	//	j.state = StateRunning
	//	var err error
	//
	//	err = j.handler(ctxJob)
	//
	//	if err != nil {
	//		j.state = StateFailed
	//		ch <- err
	//		return
	//	}
	//
	//	j.state = StateCompleted
	//	ch <- err
	//}()
	//
	////for {
	////	select {
	////	case <-j.stopChan:
	////		break
	////	default:
	////		fmt.Println("Hello world")
	////	}
	////}
	//
	////go func() {
	////	for {}
	////}()
	//
	//select {
	//case err := <-ch:
	//	doneFunc()
	//	return err
	//case <-j.stopChan:
	//	doneFunc()
	//	return nil
	//}

	//return <-ch
}

func (j *job) Retry(ctx context.Context) error {
	//if j.retryIndex == len(j.config.Retries)-1 {
	//	return nil
	//}

	j.retryIndex += 1
	time.Sleep(j.config.Retries[j.retryIndex])

	//j.state = StateRunning
	err := j.Execute(ctx)

	if err == nil {
		j.state = StateCompleted
		return nil
	}

	if j.retryIndex == len(j.config.Retries)-1 { // If Retry index is already the last one of the list
		j.state = StateRetryFailed
		return err
	}

	j.state = StateFailed // If Retry index is not yet the last and still to be continued...
	return err
}

func (j *job) State() JobState { return j.state }
func (j *job) RetryIndex() int { return j.retryIndex }

func (j *job) SetRetryDurations(times []time.Duration) {
	if len(times) == 0 {
		return
	}

	j.config.Retries = times
}

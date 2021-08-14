package ratelimiter

import (
	"context"
	"sync"
	"time"
)

type RateLimit struct {
	// Per second allowed
	Per int64
	// Runner represents the function will perform
	// each item at time
	Runner Runner
	// ResponseChan represents the data returned by the Runner
	// after be processed
	ResponseChan chan interface{}
	// ErrChan is a chan will receive item and error
	// when the Runner fail.
	ErrChan chan Err
	// queueChan will store item in pending state
	queueChan chan *ItemContextual
	// internal usage. Tu stop a routine use RoutineStop()
	// function
	running  bool
	stopChan chan bool
	mutex    *sync.Mutex
}

type ItemContextual struct {
	Ctx  context.Context
	Item interface{}
}

type Err struct {
	Error error
	Item  *ItemContextual
}

// Runner function is used to resole function
type Runner func(ctx context.Context, item interface{}) (interface{}, error)

func New(limitPerSecond int, runner Runner) *RateLimit {
	return &RateLimit{
		mutex:        &sync.Mutex{},
		Per:          int64(limitPerSecond),
		Runner:       runner,
		ResponseChan: make(chan interface{}),
		running:      false,
		stopChan:     make(chan bool),
		ErrChan:      make(chan Err),
		queueChan:    make(chan *ItemContextual),
	}
}

func (r *RateLimit) Push(ctx context.Context, item interface{}) {
	go func() {
		r.mutex.Lock()
		r.queueChan <- &ItemContextual{ctx, item}
		r.mutex.Unlock()
	}()
}

func (r *RateLimit) RoutineStart() {
	if r.running {
		return
	}

	r.running = true
	go func() {
	Loop:
		for {
			select {
			case <-r.stopChan:
				break Loop
			case msg := <-r.queueChan:
				result, err := r.Runner(msg.Ctx, msg.Item)
				if err != nil {
					r.ErrChan <- Err{Error: err, Item: msg}
				}
				r.ResponseChan <- result
				time.Sleep(time.Duration((1 / r.Per) * int64(time.Second)))
			}
		}
	}()
}

func (r *RateLimit) RoutineStop() {
	r.stopChan <- true
	r.running = false
}

func (r *RateLimit) QueueSize() int {
	return len(r.queueChan)
}

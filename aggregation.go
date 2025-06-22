package main

import (
	"fmt"
	"sync"
	"time"
)

func DataAggregation() {
	start := time.Now()
	username := fetchUser()
	respch := make(chan any, 2) // buffered channel

	// ch := make(chan int, 3)     // Buffered channel with capacity 3
	// ch <- 1 // doesn't block
	// ch <- 2 // doesn't block
	// ch <- 3 // doesn't block
	// ch <- 4 // blocks if no receiver has read
	// An unbuffered channel requires:
	// Sender to wait for a receiver: ch <- val blocks until a receiver reads from the channel.
	// Receiver to wait for a sender: <-ch blocks until a sender writes to the channel.
	// This leads to strict synchronization between goroutines.

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go fetchUserLikes(username, respch, wg)
	go fetchUserMatch(username, respch, wg)
	// close(respch) // cannot close it now, need to synchronize -> sync waitgroup

	wg.Wait() // block until we have 2 wg done calls
	close(respch)

	for res := range respch {
		// will keep on ranging until deadlock if we do not close the channel
		fmt.Println(res)
	}

	// fmt.Printf("userName:%s likes:%d match:%s \n", username, likes, match)
	fmt.Println("took:", time.Since(start))
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)
	return "bot"
}

func fetchUserLikes(username string, respch chan any, wg *sync.WaitGroup) {
	fmt.Println("fetching user likes:", username)
	time.Sleep(time.Millisecond * 150)
	respch <- 10
	wg.Done()
}

func fetchUserMatch(username string, respch chan any, wg *sync.WaitGroup) {
	fmt.Println("fetching user match:", username)
	time.Sleep(time.Millisecond * 100)
	respch <- "ai"
	wg.Done()
}

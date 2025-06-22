package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// func TestRaceConditions() {
// 	var state int

// 	for i := 1; i < 10; i++ {
// 		state += i
// 	}
// }

// channels are cooler though
func TestRaceConditions() {
	var state int
	var mut sync.RWMutex
	for i := 1; i < 10; i++ {
		go func(i int) {
			mut.Lock() // locking to avoid races
			state += i
			mut.Unlock()
		}(i)
	}
	fmt.Println(state)

	// atomic value
	var atomicState int32
	for i := 1; i < 10; i++ {
		go func(i int) {
			atomic.AddInt32(&atomicState, int32(i))
		}(i)
	}

}

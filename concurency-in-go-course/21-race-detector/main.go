package main

import (
	"fmt"
	"math/rand"
	"time"
)

// TODO: identify the data race and fix the issue

func main() {
	start := time.Now()
	var t *time.Timer

	ch := make(chan bool) // new code

	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		// t.Reset(randomDuration()) // old code
		ch <- true // new code
	})

	// time.Sleep(5 * time.Second) // old code
	// new code
	for time.Since(start) < 5*time.Second {
		<-ch
		t.Reset(randomDuration())
	}

}

// returns random duration between 0-1 seconds
func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

//-------------
// The variable `t` is accessed by two goroutines:
// (main goroutine) -> t <- (time.AfterFunc gouroutine)
//-------------
// (working condition)
// main goroutine...
// t = time.AfterFunc() // return a timer

// AfterFunc goroutine
// t.Reset() // timer reset
// ------------
// (race condition- random duration is very small)
// AfterFunc goroutine
// t.Reset() // t = nil

// main goroutine...
// t = time.AfterFunc()
//-------------

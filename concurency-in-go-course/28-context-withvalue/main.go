package main

import (
	"context"
	"fmt"
)

type database map[string]bool
type userIDKeyType string

var db database = database{
	"bob": true,
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	processRequest(ctx, "bob")
}

func processRequest(ctx context.Context, userid string) {
	// TODO: send userID information to checkMembership through ceontext for map lookup.
	vctx := context.WithValue(ctx, userIDKeyType("userIDKey"), "bob")

	ch := checkMembership(vctx)
	status := <-ch
	fmt.Printf("membership status of userid : %s : %v\n", userid, status)
}

// checkMembership - takes context as input.
// extracts the user id infiormation from context.
// spins a goroutine to do map lookup
// sends the result on the returned chanel.
func checkMembership(ctx context.Context) <-chan bool {
	ch := make(chan bool)
	go func() {
		defer close(ch)
		// do some database lookup
		userid := ctx.Value(userIDKeyType("userIDKey")).(string)
		status := db[userid]
		ch <- status
	}()
	return ch
}

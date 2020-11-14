package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "userID", 12345)
	ctx = context.WithValue(ctx, 1, "admin")

	if v := ctx.Value("userID"); v != nil {
		fmt.Printf("userID: %d\n", v)
	} else {
		fmt.Println("no value associated with that key")
	}

	fmt.Printf("type: %s\n", ctx.Value(1).(string))

	if v := ctx.Value("userType"); v != nil {
		fmt.Printf("userType: %d\n", v)
	} else {
		fmt.Println("no value associated with the key 'userType'")
	}
}

package main

import (
	"context"
	"fmt"

	"github.com/ck3g/SomeDaysOfGo/exploring-the-go-course/06-context/ex-02/session"
)

func main() {
	ctx := session.SetUserID(context.Background(), 123)
	fmt.Printf("userID: %d\n", session.GetUserID(ctx))
	fmt.Printf("isAdmin? %t\n", session.IsAdmin(ctx))

	ctx = session.SetAdminAccess(ctx, true)
	fmt.Printf("isAdmin? %t\n", session.IsAdmin(ctx))
}

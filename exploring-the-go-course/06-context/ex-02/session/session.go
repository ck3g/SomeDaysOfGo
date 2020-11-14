package session

import "context"

type stringKey string
type intKey int

var userID stringKey
var admin intKey

func SetUserID(ctx context.Context, uID int) context.Context {
	return context.WithValue(ctx, userID, uID)
}

func SetAdminAccess(ctx context.Context, isAdmin bool) context.Context {
	return context.WithValue(ctx, admin, isAdmin)
}

func GetUserID(ctx context.Context) int {
	if v := ctx.Value(userID); v != nil {
		return v.(int)
	}

	return 0
}

func IsAdmin(ctx context.Context) bool {
	if v := ctx.Value(admin); v != nil {
		return v.(bool)
	}

	return false
}

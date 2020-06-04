package mock

import "github.com/ck3g/SomeDaysOfGo/build-your-own-x/wtf-dial"

// Authenticator represents a service fo authenticating users.
type Authenticator struct {
	AuthenticateFn      func(token string) (*wtf.User, error)
	AuthenticateInvoked bool
}

// Authenticate authenticates users by token.
func (a *Authenticator) Authenticate(token string) (*wtf.User, error) {
	a.AuthenticateInvoked = true
	return a.AuthenticateFn(token)
}

// DefaultUser is the user authenticated by DefaultAuthenticator.
var DefaultUser = &wtf.User{ID: 503}

// DefaultAuthenticator returns an authenticator that returns the default user.
func DefaultAuthenticator() Authenticator {
	return Authenticator{
		AuthenticateFn: func(token string) (*wtf.User, error) { return DefaultUser, nil },
	}
}

// DialService represents a service for managing Dials.
type DialService struct {
	DialFn      func(id wtf.DialID) (*wtf.Dial, error)
	DialInvoked bool

	CreateDialFn      func(dial *wtf.Dial) error
	CreateDealInvoked bool

	SetLevelFn      func(id wtf.DialID, level float64) error
	SetLevelInvoked bool
}

// Dial retrieves dial by ID.
func (s *DialService) Dial(id wtf.DialID) (*wtf.Dial, error) {
	s.DialInvoked = true
	return s.DialFn(id)
}

// CreateDial creates a new dial.
func (s *DialService) CreateDial(dial *wtf.Dial) error {
	s.CreateDealInvoked = true
	return s.CreateDialFn(dial)
}

// SetLevel sets the current WTF level for a user.
func (s *DialService) SetLevel(id wtf.DialID, level float64) error {
	s.SetLevelInvoked = true
	return s.SetLevelFn(id, level)
}

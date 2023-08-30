package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen  = 6
)

var (
	ErrFirstNameTooShort = fmt.Sprintf("firstName length should be at least %d characters", minFirstNameLen)
	ErrLastNameTooShort  = fmt.Sprintf("lastName length should be at least %d characters", minLastNameLen)
	ErrPasswordTooShort  = fmt.Sprintf("password length should be at least %d characters", minPasswordLen)
	ErrEmailNotValid     = "email is not valid"
)

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params CreateUserParams) Validate() []string {
	errors := []string{}

	if len(params.FirstName) < minFirstNameLen {
		errors = append(errors, ErrFirstNameTooShort)
	}
	if len(params.LastName) < minLastNameLen {
		errors = append(errors, ErrLastNameTooShort)
	}
	if len(params.Password) < minPasswordLen {
		errors = append(errors, ErrPasswordTooShort)
	}
	if !isValidEmail(params.Email) {
		errors = append(errors, ErrEmailNotValid)
	}

	return errors
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"encryptedPassword" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encpw),
	}, nil
}

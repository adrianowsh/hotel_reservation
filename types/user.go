package types

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost         = 12
	minFirstNameLength = 2
	minLastNameLength  = 2
	maxFirstNameLength = 50
	maxLastNameLength  = 50
	minPasswordLength  = 8
	validAge           = 0
	emailRegex         = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
)

type CreateUserParam struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Password  string `json:"password"`
}

type User struct {
	ID           string `bson:"_id,omitempty" json:"id,omitempty"`
	Email        string `bson:"email" json:"email"`
	FirstName    string `bson:"first_name" json:"first_name"`
	LastName     string `bson:"last_name" json:"last_name"`
	Age          int    `bson:"age" json:"age"`
	PasswordHash string `bson:"password_hash" json:"-"`
}

type ErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

type UsersResponse struct {
	Data []*User `json:"data"`
}

type UserResponse struct {
	Data *User `json:"data"`
}

func (param CreateUserParam) Validate() *ErrorResponse {
	errors := map[string]string{}

	if len(param.FirstName) < minFirstNameLength || len(param.FirstName) > maxFirstNameLength {
		errors["first_name"] = fmt.Sprintf("first name must be between %d and %d characters", minFirstNameLength, maxFirstNameLength)
	}
	if len(param.LastName) < minLastNameLength || len(param.LastName) > maxLastNameLength {
		errors["last_name"] = fmt.Sprintf("last name must be between %d and %d characters", minLastNameLength, maxLastNameLength)
	}
	if len(param.Password) < minPasswordLength {
		errors["password"] = fmt.Sprintf("password must be at least %d characters", minPasswordLength)
	}
	if len(param.FirstName) > maxFirstNameLength {
		errors["first_name"] = fmt.Sprintf("first name must be between %d and %d characters", minFirstNameLength, maxFirstNameLength)
	}
	if len(param.LastName) > maxLastNameLength {
		errors["last_name"] = fmt.Sprintf("last name must be between %d and %d characters", minLastNameLength, maxLastNameLength)
	}
	if param.Age < validAge {
		errors["age"] = fmt.Sprintf("age must be at least %d", validAge)
	}
	if matched, _ := regexp.MatchString(emailRegex, param.Email); !matched {
		errors["email"] = "invalid email format"
	}
	if len(errors) > 0 {
		return &ErrorResponse{Errors: errors}
	}
	return nil
}

func NewUserFromParam(param *CreateUserParam) (*User, error) {
	nencpw, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcryptCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Email:        param.Email,
		FirstName:    param.FirstName,
		LastName:     param.LastName,
		Age:          param.Age,
		PasswordHash: string(nencpw),
	}, nil
}

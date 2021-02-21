package user

import (
	errors "github.com/avinashb98/bookstore-users-api/utils/error"
	"strings"
)

type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	DateUpdate  string `json:"date_updated"`
}

func (user *User) Validate() *errors.RestErr {
	email := strings.TrimSpace(strings.ToLower(user.Email))
	if email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	return nil
}

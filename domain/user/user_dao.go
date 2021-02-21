package user

import (
	"fmt"
	"github.com/avinashb98/bookstore-users-api/utils/dates"
	"github.com/avinashb98/bookstore-users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}

	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.ID]
	if current != nil {

		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}
	now := dates.GetNowString()
	user.DateCreated = now
	user.DateUpdate = now
	usersDB[user.ID] = user
	return nil
}

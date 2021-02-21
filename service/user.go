package service

import (
	entity "github.com/avinashb98/bookstore-users-api/domain/user"
	errors "github.com/avinashb98/bookstore-users-api/utils/error"
)

func CreateUser(user entity.User) (*entity.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userID int64) (*entity.User, *errors.RestErr) {
	result := entity.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return &result, nil
}

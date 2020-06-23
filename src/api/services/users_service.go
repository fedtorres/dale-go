package services

import (
	"github.com/fedtorres/dale-go/src/api/domain"
	"github.com/fedtorres/dale-go/src/api/utils"
)

type usersService struct{}

var (
	UsersService usersService
)

func (u *usersService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

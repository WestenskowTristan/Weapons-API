package users

import "time"

type IUsersService interface {
	Get(userId string) (*User, error)
}

type UsersService struct {}

func InitUsersService() IUsersService {
	return &UsersService{}
}

func (service *UsersService) Get(userId string) (*User, error) {
	user := User{
		Id: userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FirstName: "Tristan",
		LastName: "Westenskow",
		Email: "joemama521@gmail.com",
		Password: "butts",
	}
	return &user, nil
}
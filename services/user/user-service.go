package services

import (
	userCliente "user-backend/clients/user"
	dto "user-backend/dtos/user"
	userModel "user-backend/models/user"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (dto.UserDto, error)
	GetUsers() (dto.UsersDto, error)
	InsertUser(dto.InsertUserDto) error
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (dto.UserDto, error) {

	user, err := userCliente.GetUserById(id)
	var userDto dto.UserDto

	if err != nil {
		return userDto, err
	}
	if user.UserID == 0 {
		return userDto, nil
	}
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Id = user.UserID
	userDto.Email = user.Email
	userDto.Password = user.Pwd
	return userDto, nil
}

func (s *userService) GetUsers() (dto.UsersDto, error) {
	users, err := userCliente.GetUsers()
	var usersDto dto.UsersDto

	if err != nil {
		return usersDto, err
	}

	for _, user := range users {
		var userDto dto.UserDto
		userDto.Name = user.Name
		userDto.LastName = user.LastName
		userDto.UserName = user.UserName
		userDto.Id = user.UserID
		userDto.Email = user.Email
		userDto.Password = user.Pwd
		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) InsertUser(userDto dto.InsertUserDto) error {

	newUser := userModel.User{
		Name:     userDto.Name,
		LastName: userDto.LastName,
		UserName: userDto.UserName,
		Email:    userDto.Email,
		Pwd:      userDto.Password,
	}

	err := userCliente.InsertUser(newUser)
	if err != nil {
		return err
	}

	return nil
}
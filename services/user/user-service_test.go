package services_test

import (
	"testing"
	dto "user-backend/dtos/user"
)

type userServiceMock struct{}

type userServiceMockInterface interface {
	GetUserById(id int) (dto.UserDto, error)
	GetUsers() (dto.UsersDto, error)
	InsertUser(dto.InsertUserDto) error
}

var (
	UserServiceMock userServiceMockInterface
)

func init() {
	UserServiceMock = &userServiceMock{}
}

func (s *userServiceMock) GetUserById(id int) (dto.UserDto, error) {
	userMocked := dto.UserDto{
		Id:       1,
		Name:     "John",
		LastName: "Doe",
		UserName: "johndoe",
		Email:    "john@example.com",
		Password: "password",
	}
	if id == 1 {
		return userMocked, nil
	}

	return dto.UserDto{}, nil
}

func (s *userServiceMock) GetUsers() (dto.UsersDto, error) {
	var user1, user2 dto.UserDto
	var users dto.UsersDto

	users = append(users, user1, user2)

	return users, nil
}

func (s *userServiceMock) InsertUser(userDto dto.InsertUserDto) error {
	// Implementación de ejemplo para método InsertUser
	return nil
}

func TestUserServiceMock_GetUserById(t *testing.T) {
	mock := &userServiceMock{}

	// Llamada al método GetUserById de la interfaz userServiceMockInterface
	user, err := mock.GetUserById(1)

	// Verificación de errores
	if err != nil {
		t.Errorf("Error en GetUserById: %v", err)
	}

	// Verificación de valores devueltos
	expectedUser := dto.UserDto{
		Id:       1,
		Name:     "John",
		LastName: "Doe",
		UserName: "johndoe",
		Email:    "john@example.com",
		Password: "password",
	}

	if user != expectedUser {
		t.Errorf("Resultado incorrecto. Se esperaba %v pero se obtuvo %v", expectedUser, user)
	}
}

func TestUserServiceMock_GetUsers(t *testing.T) {
	mock := &userServiceMock{}

	// Llamada al método GetUsers de la interfaz userServiceMockInterface
	users, err := mock.GetUsers()

	// Verificación de errores
	if err != nil {
		t.Errorf("Error en GetUsers: %v", err)
	}

	// Verificación de cantidad de usuarios obtenidos
	if len(users) != 2 {
		t.Errorf("Número incorrecto de usuarios. Se esperaban 1, se obtuvieron %d", len(users))
	}
}

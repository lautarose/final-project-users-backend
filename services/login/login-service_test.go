package services_test

import (
	"errors"
	"testing"
	userModel "user-backend/models/user"
)

// Mock del cliente userCliente (dummy)
type dummyUserCliente struct{}

func (d *dummyUserCliente) GetUserByUsername(username string) (userModel.User, error) {
	// Simular un usuario de prueba
	if username == "testuser" {
		return userModel.User{
			UserID:   1,
			UserName: "testuser",
			Pwd:      "password123",
		}, nil
	}
	return userModel.User{}, errors.New("usuario no encontrado")
}

func TestUserCliente_GetUserByUsername(t *testing.T) {
	// Crear instancia del mock del cliente userCliente (dummy)
	mockClient := &dummyUserCliente{}

	// Ejecutar el método GetUserByUsername del cliente dummy
	foundUser, err := mockClient.GetUserByUsername("testuser")

	// Verificar si no hay error en la obtención del usuario
	if err != nil {
		t.Errorf("Se esperaba un error nulo pero se obtuvo: %v", err)
	}

	// Verificar si se obtuvo el usuario correctamente
	if foundUser.UserID != 1 || foundUser.UserName != "testuser" || foundUser.Pwd != "password123" {
		t.Error("Los datos del usuario obtenidos no son los esperados")
	}
}

package app

import (
	loginController "user-backend/controllers/login"
	userController "user-backend/controllers/user"
)

// MapUrls maps the urls
func MapUrls() {

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/users", userController.GetUsers)
	router.POST("/user", userController.InsertUser)

	// Login Mapping
	router.POST("/user/login", loginController.Login)

}

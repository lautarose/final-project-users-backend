package app

import (
	userController "user-backend/controllers/user"
)

// MapUrls maps the urls
func MapUrls() {

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/users", userController.GetUsers)
	router.POST("/user", userController.InsertUser)

}
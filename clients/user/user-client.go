package clients

import (
	model "user-backend/models/user"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUserById(id int) (model.User, error) {
	var user model.User

	err := Db.Where("user_id = ?", id).First(&user).Error

	if err != nil {
		log.Println(err)
		return user, err
	}

	log.Debug("User: ", user)

	return user, nil
}

func GetUserByUsername(username string) (model.User, error) {
	var user model.User

	err := Db.Where("user_name = ?", username).First(&user).Error

	if err != nil {
		log.Println(err)
		return user, err
	}
	log.Debug("User: ", user)

	return user, nil
}

func GetUsers() (model.Users, error) {

	var users model.Users

	err := Db.Find(&users).Error

	if err != nil {
		log.Println(err)
		return users, err
	}

	log.Debug("Users: ", users)

	return users, nil
}

func InsertUser(newUser model.User) error {
    err := Db.Create(&newUser).Error
    if err != nil {
        log.Println(err)
        return err
    }

    log.Debug("User inserted: ", newUser)

    return nil
}
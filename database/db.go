package database

import (
	userClient "user-backend/clients/user"
	userModel "user-backend/models/user"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "users-db" //variable de entorno para nombre de la base de datos
	DBUser := "usersdb"  //variable de entorno para el usuario de la base de datos
	//DBPass := ""
	DBPass := "password" //variable de entorno para la pass de la base de datos
	DBHost := "34.95.218.145"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+")/"+DBName+"?charset=utf8&parseTime=True")

	db.LogMode(true)

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	userClient.Db = db
}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&userModel.User{})

	log.Info("Finishing Migration Database Tables")
}

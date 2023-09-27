package model

import (
	"OrmGo/config"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitModel(config config.ProgramConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println(connectionString)
		logrus.Error("Model : cannot connect to database, ", err.Error())
		return nil

	}
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Users{})
	db.AutoMigrate(&Books{})
	db.AutoMigrate(&Blog{})
	db.Model(&Users{}).Association("Blogs").Append(&Blog{})
}

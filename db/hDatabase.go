package db

import (
	"GO_SIMPLECRUD/models"
	"log"
	"os"

	yaml "gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(path string) *gorm.DB {

	//protocol://username:password@host:port/database

	conf := &dbConfig{}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(&conf); err != nil {
		panic(err)
	}

	dbURL := conf.Protocol + "://" + conf.Username + ":" + conf.Password + "@" + conf.Host + ":" + conf.Port + "/" + conf.DB

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Plants{})
	db.AutoMigrate(&models.Rooms{})
	db.AutoMigrate(&models.Species{})
	db.AutoMigrate(&models.WateringDays{})

	return db
}

func NewDB() *gorm.DB {

	return InitDB("db/config.yml")

}

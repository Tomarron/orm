package orm

import (
	"log"

	base "github.com/jinzhu/gorm"
)

type Model base.Model

type Settings struct {
	Driver string
	Path   string
}

var DB *base.DB
var settings Settings

func Initialize(driver, path string) {
	log.Println("Initialize Gorm")
	if driver == "sqlite" {
		driver = "sqlite3"
	}
	settings = Settings{
		Driver: driver,
		Path:   path,
	}
}

func connect() {
	d, err := base.Open(settings.Driver, settings.Path)
	if err != nil {
		panic(err)
	}

	DB = d
}

func Get() *base.DB {
	if DB != nil {
		DB.Close()
	}
	connect()
	return DB
}

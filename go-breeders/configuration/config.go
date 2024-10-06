package configuration

import (
	"database/sql"
	"go-breeders/models"
	"sync"
)

type Application struct {
	Models *models.Models
}

var instance *Application
var once sync.Once
var db *sql.DB

func New(pool *sql.DB) *Application {
	db = pool
	return GetInstance()
}

func GetInstance() *Application {
	// Do call the function if and only if do is being called for the first time for this instance of once
	once.Do(func() {
		instance = &Application{
			Models: models.New(db),
		}
	})
	return instance
}

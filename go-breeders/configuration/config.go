package configuration

import (
	"database/sql"
	"go-breeders/adapters"
	"go-breeders/models"
	"sync"
)

type Application struct {
	Models *models.Models
	CatService  *adapters. RemoteService

}

var instance *Application
var once sync.Once
var db *sql.DB
var catSevice *adapters.RemoteService

func New(pool *sql.DB, cs *adapters.RemoteService) *Application {
	db = pool
	catSevice = cs
	return GetInstance()
}

func GetInstance() *Application {
	// Do call the function if and only if do is being called for the first time for this instance of once
	once.Do(func() {
		instance = &Application{
			Models: models.New(db),
			CatService: catSevice,
		}
	})
	return instance
}

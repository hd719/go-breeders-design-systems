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
var once sync.Once // Allow us to create our singleton
var db *sql.DB

func New(pool *sql.DB) *Application {
	db = pool
	return GetInstance()
}

func GetInstance() *Application {
	// Do takes a function and call it exactly 1 time when the application is started
	once.Do(func() {
		instance = &Application{
			Models: models.New(db), // hooking up the models with the database connection
		}
	})

	return instance
}

// In our example:
// Let's say we want to open up another database connection in some other package, we can call models.New
// Note: when we call models.New() we are grabbing certain number of database connections that are available from the database server - eventually those will get exhausted if keep calling models.New() in our other packages
// Instead, once we call configuration.New() -> we wont open any NEW connections to the database because during start up a connection already exists!
// Future note: Re-watch video 51...

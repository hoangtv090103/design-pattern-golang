package models

import "database/sql"

// Repository is the database repository. Anything that implements
// This interface must implement all the methods included here.
type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
	GetBreedByName(b string) (*DogBreed, error)
}

// mysqlRepository is a simple wrapper for the *sql.DB type.
// This is used to return a MySQL/MariaDB repository
type mysqlRepository struct {
	DB *sql.DB
}

// newMysqlRepository is a convenience factory method to return a new mysqlRepository.
func newMysqlRepository(conn *sql.DB) Repository {
	return &mysqlRepository{
		DB: conn,
	}
}

type testRepository struct {
	DB *sql.DB // never going to actually put anything in that field
}

func newTestMysqlRepository(_ *sql.DB) Repository {
	return &testRepository{
		DB: nil,
	}
}

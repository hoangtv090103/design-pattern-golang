package main

// This is where we setup our test environment
import (
	"go-breeders/configuration"
	"go-breeders/models"
	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
    
    testBackend := &TestBackend{}
    testAdapter := &RemoteService{
        Remote: testBackend,
    }

	// test app
	testApp = application{
		App: configuration.New(nil),
		catService: testAdapter,
	}

	os.Exit(m.Run())
}

type TestBackend struct {}

func (tb *TestBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
    breeds := []*models.CatBreed{
        {
            ID: 1,
            Breed: "Tomcat",
            Details: "Some details",
        },
    }
    
    return breeds, nil
}

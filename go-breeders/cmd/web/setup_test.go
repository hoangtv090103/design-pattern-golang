package main

// This is where we setup our test environment
import (
	"go-breeders/adapters"
	"go-breeders/configuration"
	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
    
    testBackend := &adapters.TestBackend{}
    testAdapter := &adapters.RemoteService{
        Remote: testBackend,
    }

	// test app
	testApp = application{
		App: configuration.New(nil, testAdapter),
	}

	os.Exit(m.Run())
}

package main

import (
	"fmt"
	"go-breeders/pets"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
	// Get a new pet of type dog
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("dog"))
}

func (app *application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
	// Get a new pet of type Cat
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("cat"))
}

func (app *application) CreateDogFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dog, err := pets.NewPetFromAbstractFactory("dog")

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, dog)
}

func (app *application) CreateCatFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	cat, err := pets.NewPetFromAbstractFactory("cat")

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, cat)
}

func (app *application) GetAllDogBreedsJSON(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dogBreeds, err := app.App.Models.DogBreed.All()

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, dogBreeds)

}

func (app *application) CreateDogWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	// create a dog using the builder pattern
	p, err := pets.NewPetBuilder().
		SetSpecies("dog").
		SetBreed("mixed breed").
		SetWeight(15).
		SetDescription("A mixed breed of unknown origin. Probably has some German Shepherd heritage").
		SetColor("Black and White").
		SetAge(3).
		SetAgeEstimated(true).
		Build()

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, p)

}

func (app *application) CreateCatWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	// Create a cat using the builder pattern
	p, err := pets.NewPetBuilder().
		SetSpecies("cat").
		SetBreed("fellis silvestris catus").
		SetWeight(4).
		SetDescription("A beautiful house cat.").
		SetGeographicOrigin("Canada").
		SetColor("Calico").
		SetAge(1).
		SetAgeEstimated(true).
		Build()

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, p)
}

func (app *application) GetAllCatBreeds(w http.ResponseWriter, r *http.Request) {
    var t toolbox.Tools

    catBreeds, err := app.App.CatService.GetAllBreeds()

    if err != nil {
        _ = t.ErrorJSON(w, err, http.StatusBadRequest)
        return
    }

    _ = t.WriteJSON(w, http.StatusOK, catBreeds)
}

func (app *application) AnimalFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
    // Setup toolbox
    var t toolbox.Tools
    
    // Get species from URL itself
    species := chi.URLParam(r, "species")
    
    // Get breed from the URL
    b := chi.URLParam(r, "breed")
    
    // In case breed is 2 or more words, hello world -> hello%20wo
    breed, _ := url.QueryUnescape(b)
    
    fmt.Println("Species:", species, "Breed:", breed)
    
    // Create a pet from abstract factory
    pet, err := pets.NewPetWithBreedFromAbstractFactory(species, breed)
    if err != nil {
        _ = t.ErrorJSON(w, err, http.StatusBadRequest)
        return
    }
    // Write the result as JSON
    _ = t.WriteJSON(w, http.StatusOK, pet)
}
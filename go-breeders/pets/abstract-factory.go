package pets

import (
	"errors"
	"fmt"
	"go-breeders/configuration"
	"go-breeders/models"
)

type IAnimal interface {
	Show() string
}

type DogFromFactory struct {
	Pet *models.Dog
}

func (dff *DogFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", dff.Pet.Breed.Breed)
}

type CatFromFactory struct {
	Pet *models.Cat
}

func (cff *CatFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", cff.Pet.Breed.Breed)
}

type IPetFactory interface {
	// returns an object that satisfies the IAnimal interface.
	newPet() IAnimal
	newPetWithBreed(breed string) IAnimal
}

type DogAbstractFactory struct {
}

func (df *DogAbstractFactory) newPet() IAnimal {
	return &DogFromFactory{
		Pet: &models.Dog{},
	}
}

func (df *DogAbstractFactory) newPetWithBreeds(b string) IAnimal {
    app := configuration.GetInstance()
    breed, _ := app.Models.DogBreed.GetBreedByName(b)
    return &DogFromFactory{
        Pet: &models.Dog{
            Breed: *breed,
        },
    }
}

type CatAbstractFactory struct {
}

func (cf *CatAbstractFactory) newPet() IAnimal {
	return &CatFromFactory{
		Pet: &models.Cat{},
	}
}

func (cf *CatAbstractFactory) newPetWithBreeds(breed string) IAnimal {
    // Get breed for cat

    // app := configuration.GetInstance()
    // breed, _ := app.Models.DogBreed.GetBreedByName(breed)
    return &CatFromFactory{
        Pet: &models.Cat{
            // Breed: breed,
        },
    }
}

func NewPetFromAbstractFactory(species string) (IAnimal, error) {
	switch species {
	case "dog":
		var dogFactory DogAbstractFactory
		dog := dogFactory.newPet()
		return dog, nil
	case "cat":
		var catFactory CatAbstractFactory
		cat := catFactory.newPet()
		return cat, nil
	default:
		return nil, errors.New("invaild specie supplied")
	}
}

// NewPetWithBreedFromAbstractFactory is the abstract factory method
func NewPetWithBreedFromAbstractFactory(
	species, breed string,
) (IAnimal, error) {
	switch species {
	case "dog":
		// return a dog with breed embedded
		var dogFactory DogAbstractFactory
		dog := dogFactory.newPetWithBreed(breed)
		return dog, nil
		
	case "cat":
		// return a cat with breed embedded
		var catFactory CatAbstractFactory
		cat := catFactory.newPet()
		return cat, nil
	default:
		return nil, errors.New("invalid species supplied")
	}
}

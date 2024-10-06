package main

import "fmt"

// Animal is the type for our abstract factory.
type Animal interface {
	Says()
	LikesWater() bool
}

// Dog is the a concrete factory for dogs.
type Dog struct{}

// Implement the abstract factory for dogs
func (d *Dog) Says() {
	fmt.Println("Woof")
}

func (d *Dog) LikesWater() bool {
	return true
}

// Create a concrete factory for cats
type Cat struct{}

// Implement the abstract factory for cats
func (c *Cat) Says() {
	fmt.Println("Meow")
}

func (c *Cat) LikesWater() bool {
	return false
}

type AnimalFactory interface { // Only to make it readable
	New() Animal
}

type DogFactory struct {
}

func (df *DogFactory) New() Animal {
	return &Dog{}
}

type CatFactory struct {
}

func (cf *CatFactory) New() Animal {
	return &Cat{}
}

func main() {
	// Create one each of a DogFactory and a CatFactory

	DogFactory := DogFactory{}
	CatFactory := CatFactory{}

	// Cal the New method to create a dog and a cat
	dog := DogFactory.New()
	cat := CatFactory.New()

	dog.Says()
	cat.Says()

	fmt.Println("A dog likes water:", dog.LikesWater())
	fmt.Println("A cat likes water:", cat.LikesWater())
}

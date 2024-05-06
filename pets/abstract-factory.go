package pets

import (
	"errors"
	"fmt"
	"go-breeders/models"
)

type AnimalInterface interface {
	Show() string
}

type DogFromFactory struct {
	Pet *models.Dog
}

func (dff *DogFromFactory) Show() string {
	fmt.Println("DogFromFactory:", dff)
	fmt.Println("Pet:", dff.Pet)
	fmt.Println("Breed:", dff.Pet.Breed)
	formattedString := fmt.Sprintf("This animal is a %s", dff.Pet.Breed.Breed)
	fmt.Println(formattedString)
	return fmt.Sprintf("This animal is a %s", dff.Pet.Breed.Breed)
}

type CatFromFactory struct {
	Pet *models.Cat
}

func (cff *CatFromFactory) Show() string {
	formattedString := fmt.Sprintf("This animal is a %s", cff.Pet.Breed.Breed)
	fmt.Println(formattedString)
	return fmt.Sprintf("This animal is a %s", cff.Pet.Breed.Breed)
}

type PetFactoryInterface interface {
	newPet() AnimalInterface
}

type DogAbstractFactory struct{}

func (df *DogAbstractFactory) newPet() AnimalInterface {
	return &DogFromFactory{
		Pet: &models.Dog{},
	}
}

type CatAbstractFactory struct{}

func (df *CatAbstractFactory) newPet() AnimalInterface {
	return &CatFromFactory{
		Pet: &models.Cat{},
	}
}

func NewPetFromAbstractFactory(species string) (AnimalInterface, error) {
	switch species {
	case "dog":
		var dogFactory DogAbstractFactory
		dog := dogFactory.newPet()
		dog.Show()
		return dog, nil
	case "cat":
		var catFactory CatAbstractFactory
		cat := catFactory.newPet()
		cat.Show()
		return cat, nil
	default:
		return nil, errors.New("invalid species supplied")
	}
}

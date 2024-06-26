package pets

import (
	"errors"
	"fmt"
	"go-breeders/configuration"
	"go-breeders/models"
	"log"
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
	newPetWithBreed(breed string) AnimalInterface
}

type DogAbstractFactory struct{}

func (df *DogAbstractFactory) newPet() AnimalInterface {
	return &DogFromFactory{
		Pet: &models.Dog{},
	}
}

func (df *DogAbstractFactory) newPetWithBreed(b string) AnimalInterface {
	app := configuration.GetInstance()
	breed, _ := app.Models.DogBreed.GetBreedByName(b)
	return &DogFromFactory{
		Pet: &models.Dog{
			Breed: *breed,
		},
	}
}

type CatAbstractFactory struct{}

func (df *CatAbstractFactory) newPet() AnimalInterface {
	return &CatFromFactory{
		Pet: &models.Cat{},
	}
}

func (cf *CatAbstractFactory) newPetWithBreed(b string) AnimalInterface {
	// Get breed for cat
	app := configuration.GetInstance()
	breed, err := app.CatService.Remote.GetCatBreedByName(b)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &CatFromFactory{
		Pet: &models.Cat{
			Breed: *breed,
		},
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

func NewPetWithBreedFromAbstractFactory(species, breed string) (AnimalInterface, error) {
	switch species {
	case "dog":
		// return a dog with breed embedded from our database
		var dogFactory DogAbstractFactory
		dog := dogFactory.newPetWithBreed(breed)
		return dog, nil
	case "cat":
		// return cat with a breed embedded from a remote service
		var catFactory CatAbstractFactory
		cat := catFactory.newPetWithBreed(breed)
		return cat, nil
	default:
		return nil, errors.New("invalid species supplied")
	}
}

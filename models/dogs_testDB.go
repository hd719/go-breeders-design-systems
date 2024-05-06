package models

func (models *testRepository) AllDogBreeds() ([]*DogBreed, error) {
	var breeds []*DogBreed

	return breeds, nil
}

func (m *testRepository) GetBreedByName(b string) (*DogBreed, error) {
	return nil, nil
}

func (m *testRepository) GetDogOfMonthByID(id int) (*DogOfMonth, error) {
	return nil, nil
}

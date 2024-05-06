# Design Patterns

## Factory Pattern

- Part of the builder pattern
- The simplest of all design patterns
- Create an instance of an object with the sensible default values
- Using the new keyword

Example:

```go
package main

import ("fmt", "myapp/products", "time")

func main() {
  factory := products.Product{}
  product := factory.New()
  fmt.Println("Product was created at", product.CreatedAt.UTC())

  // Not a good idea! Works for simple examples, however it is best practice to use the new keyword (there could be default values that need to be set)
  // product := products.Product{
  //   Name: "Product",
  //   CreatedAt: time.Now(),
  //   UpdatedAt: time.Now(),
  // }
}
```

```go
package products

import "time"

type Product struct {
  Name string
  CreatedAt time.Time
  UpdatedAt time.Time
}

func (p *Product) New() *Product  {
  product := Product{
    Name: "Product",
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }

  return &product
}
```

## Abstract Factory Pattern

```go
package main

// Animal is the type for abstract factory -> Animal is a abstract factory
type Animal interface {
  Says() string
  LikesWater() bool
}

// Dog is the concrete factory for dogs -> Dog is a concrete factory
type Dog struct {}

func (d *Dog) Says() string {
  fmt.Println("Woof")
}

func (d *Dog) LikesWater() bool {
  return true
}

// Cat is the concrete factory for dogs -> Cat is a concrete factory
type Cat struct {}

func (c *Cat) Says() string {
  fmt.Println("Meow")
}

func (c *Cat) LikesWater() bool {
  return false
}

// We will not be using this, but makes the code more readable
type AnimalFactory interface {
  New() Animal
}

type DogFactory struct {}

func (df *DogFactory) New() Animal {
  return &Dog{}
}

type CatFactory struct {}

func (cf *CatFactory) New() Animal {
  return &Cat{}
}

func main() {
  // Create one each of a DogFactory and CatFactory
  dogFactory := DogFactory{}
  catFactory := CatFactory{}

  // Call the new method to create a dog and cat
  dog := dogFactory.New()
  cat := catFactory.New()

  dog.Says() // Woof
  cat.Says() // Meow
}
```

- Create families of related objects without relying on their concrete classes using interfaces
- Decouple client code from the concrete classes (we just want the FE to call some method and they get their object), for example in our case we may want to create a dog object from the database and the cat object from some remote api service

## Repository Pattern

- The repository pattern is a design pattern that abstracts the data store from the rest of the application
- Actually an example of the Adapter pattern
- Allows change databases with ease
- Makes writing unit tests much simpler
- An intermediary layer between an applications business logic and data storage

## Singleton Pattern

- Be careful with the singleton pattern, it can be an anti-pattern (only one instance of an object is created)
- It is a global state, only use it when you need to share state across the application (in our example)

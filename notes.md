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

## Builder Pattern and Fluent Interface

- Allow us to chain methods
  - Ex.

    - ```go
      p, err := pets.NewPetBuilder().SetName("Fido").SetType("Dog").SetAge(3).Build()
      ```

  - Ex.

    - ```go
      package main

      import "fmt"

      func main() {
        address := CreateAddress().
          SetStreet("123 Main St").
          SetCity("Springfield").
          SetState("IL").
          SetZip("62701")

        fmt.Println(address)
      }

      type Address struct {
        Street string
        City string
        State string
        Zip string
      }

      func CreateAddress() *Address {
        return &Address{}
      }

      func (a *Address) SetStreet(street string) *Address {
        a.Street = street
        return a
      }

      func (a *Address) SetCity(city string) *Address {
        a.City = city
        return a
      }

      func (a *Address) SetState(state string) *Address {
        a.State = state
        return a
      }

      func (a *Address) SetZip(zip string) *Address {
        a.Zip = zip
        return a
      }
      ```

## Adapter Pattern

- Most common and use patterns
- Allows us to have different programs or parts of the same program to communicate with each other
- Ex.
  - A handler gets information from something and expects it to be in a JSON format

  ```go
  package main

  import (
    "encoding/json"
    "encoding/xml"
    "fmt"
    "io"
    "log"
    "net/http"
  )

  // ToDo is the type for the data we are working with.
  type ToDo struct {
    UserID    int    `json:"userId" xml:"userId"`
    ID        int    `json:"id" xml:"id"`
    Title     string `json:"title" xml:"title"`
    Completed bool   `json:"completed" xml:"completed"`
  }

  // DataInterface is simply our target interface, which defines all the methods that
  // any type which implements this interface must have. In our case, we only have
  // one, but you might have one for each of Create, Read, Update, Delete, and more.
  type DataInterface interface {
    GetData() (*ToDo, error)
  }

  // RemoteService is the Adaptor type. It embeds a DataInterface interface
  // (which is critical to the pattern). It is a simple wrapper for this interface.
  type RemoteService struct {
    Remote DataInterface
  }

  // CallRemoteService is the function on RemoteService which lets us
  // call any adaptor which implements the DataInterface type.
  func (rs *RemoteService) CallRemoteService() (*ToDo, error) {
    return rs.Remote.GetData()
  }

  // JSONBackend is the JSON adaptee, which needs to satisfy the DataInterface by
  // have a GetData method.
  type JSONBackend struct{}

  // GetData is necessary so that JSONBackend satisifies the DataInterface requirements.
  func (jb *JSONBackend) GetData() (*ToDo, error) {
    resp, err := http.Get("<https://jsonplaceholder.typicode.com/todos/1>")
    if err != nil {
      return nil, err
    }

    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)

    var todo ToDo
    err = json.Unmarshal(body, &todo)
    if err != nil {
      return nil, err
    }

    return &todo, nil
  }

  // XMLBackend is the XML adaptee, which needs to satisfy the DataInterface by
  // have a GetData method.
  type XMLBackend struct{}

  // GetData is necessary so that JSONBackend satisifies the DataInterface requirements.
  func (xb *XMLBackend) GetData() (*ToDo, error) {
    xmlFile := `
      <?xml version="1.0" encoding="UTF-8" ?>
        <root>
          <userId>1</userId>
          <id>1</id>
          <title>delectus aut autem</title>
          <completed>false</completed>
        </root>
    `

    var todo ToDo
    err := xml.Unmarshal([]byte(xmlFile), &todo)
    if err != nil {
      return nil, err
    }

    return &todo, nil
  }

  func main() {
    // No adapter
    todo := getRemoteData()
    fmt.Println("TODO without adapter:\t", todo.ID, todo.Title)

    // With adapter, using JSON
    jsonBackend := &JSONBackend{}
    jsonAdapter := &RemoteService{Remote: jsonBackend}
    tdFromJSON, _ := jsonAdapter.CallRemoteService()
    fmt.Println("From JSON Adapter:\t", tdFromJSON.ID, tdFromJSON.Title)

    // With adapter, using XML
    xmlBackend := &XMLBackend{}
    xmlAdapter := &RemoteService{Remote: xmlBackend}
    tdFromXML, err := xmlAdapter.CallRemoteService()
    if err != nil {
      log.Println(err)
    }
    fmt.Println("From XML Adapter:\t", tdFromXML.ID, tdFromXML.Title)
  }

  func getRemoteData() *ToDo {
    resp, err := http.Get("<https://jsonplaceholder.typicode.com/todos/1>")
    if err != nil {
      log.Fatalln(err)
    }

    defer resp.Body.Close()
    body, _ := io.ReadAll(resp.Body)

    var todo ToDo
    err = json.Unmarshal(body, &todo)
    if err != nil {
      log.Fatalln(err)
    }

    return &todo
  }

  ```

## Next Steps (Bringing it all together)

- Creating a new abstract factory method
- Creating JSON and XML adapters
- Updating our singleton with additional data
- Trying it out on the front end

## Decorator Pattern

- The decorator pattern is a structural pattern that allows us to add new behavior to objects dynamically by placing them inside special wrapper objects called decorators
- Takes an object and decorates it with additional functionality
- Simple in Go bc go uses composition instead of inheritance
- Embed a struct in another struct

## Worker Pool Pattern

- Units of work are distributed among a pool of workers which are running concurrently
- Communication is done through channels
- Useful for long running tasks or CPU bound tasks

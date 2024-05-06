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

- Create families of related objects without relying on their concrete classes

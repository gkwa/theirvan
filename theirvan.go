package theirvan

import (
	"fmt"
	"log/slog"
)

func Main() int {
	slog.Debug("theirvan", "test", true)

	main()
	return 0
}

// Product is the interface that defines the behavior of the products.
type Product interface {
	Use()
}

// ConcreteProductA is a concrete implementation of the Product interface.
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() {
	fmt.Println("Using Concrete Product A")
}

// ConcreteProductB is another concrete implementation of the Product interface.
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() {
	fmt.Println("Using Concrete Product B")
}

// Factory is the interface for creating products.
type Factory interface {
	CreateProduct() Product
}

// ConcreteFactoryA is a concrete implementation of the Factory interface for creating ConcreteProductA.
type ConcreteFactoryA struct{}

func (f *ConcreteFactoryA) CreateProduct() Product {
	return &ConcreteProductA{}
}

// ConcreteFactoryB is a concrete implementation of the Factory interface for creating ConcreteProductB.
type ConcreteFactoryB struct{}

func (f *ConcreteFactoryB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func main() {
	// Create a ConcreteFactoryA and use it to create a ConcreteProductA.
	factoryA := &ConcreteFactoryA{}
	productA := factoryA.CreateProduct()
	productA.Use()

	// Create a ConcreteFactoryB and use it to create a ConcreteProductB.
	factoryB := &ConcreteFactoryB{}
	productB := factoryB.CreateProduct()
	productB.Use()
}

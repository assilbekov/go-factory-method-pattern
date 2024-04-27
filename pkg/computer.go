package pkg

import "fmt"

type Computer interface {
	GetType() string
	PrintDetails()
}

type Laptop struct{}

func (l *Laptop) GetType() string {
	return "Laptop"
}

func (l *Laptop) PrintDetails() {
	fmt.Println("Laptop")
}

func New(typeName string) Computer {
	switch typeName {
	case "laptop":
		return &Laptop{}
	default:
		fmt.Println("Unknown type")
		return nil
	}
}

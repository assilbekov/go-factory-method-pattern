package pkg

import "fmt"

type Laptop struct{}

func (l *Laptop) GetType() string {
	return "Laptop"
}

func (l *Laptop) PrintDetails() {
	fmt.Println("Laptop")
}

func NewLaptop() Computer {
	return &Laptop{}
}

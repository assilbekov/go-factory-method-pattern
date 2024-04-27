package pkg

import "fmt"

const (
	ServerType = "server"
	ClientType = "client"
	LaptopType = "laptop"
)

type Computer interface {
	GetType() string
	PrintDetails()
}

func New(typeName string) Computer {
	switch typeName {
	case LaptopType:
		return NewLaptop()
	case ServerType:
		return NewServer()
	case ClientType:
		return NewClient()
	default:
		fmt.Println("Unknown type")
		return nil
	}
}

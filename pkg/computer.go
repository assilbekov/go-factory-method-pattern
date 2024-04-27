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

type Laptop struct{}

func (l *Laptop) GetType() string {
	return "Laptop"
}

func (l *Laptop) PrintDetails() {
	fmt.Println("Laptop")
}

type Client struct{}

func (c *Client) GetType() string {
	return "Client"
}

func (c *Client) PrintDetails() {
	fmt.Println("Client")
}

func New(typeName string) Computer {
	switch typeName {
	case LaptopType:
		return &Laptop{}
	case ServerType:
		return &Server{}
	case ClientType:
		return &Client{}
	default:
		fmt.Println("Unknown type")
		return nil
	}
}

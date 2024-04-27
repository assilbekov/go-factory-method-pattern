package pkg

import "fmt"

type Client struct{}

func (c *Client) GetType() string {
	return "Client"
}

func (c *Client) PrintDetails() {
	fmt.Println("Client")
}

func NewClient() Computer {
	return &Client{}
}

package pkg

import "fmt"

type Server struct{}

func (s *Server) GetType() string {
	return "Server"
}

func (s *Server) PrintDetails() {
	fmt.Println("Server")
}

func NewServer() Computer {
	return &Server{}
}

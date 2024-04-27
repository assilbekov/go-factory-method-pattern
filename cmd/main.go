package main

import "factory-method/pkg"

var computerTypes = []string{pkg.LaptopType, pkg.ServerType, pkg.ClientType, "unknown"}

func main() {
	for _, typeName := range computerTypes {
		computer := pkg.New(typeName)
		if computer != nil {
			computer.PrintDetails()
		}
	}
}

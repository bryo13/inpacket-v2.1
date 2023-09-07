package nic

import (
	"fmt"
	"log"
	"net"
)

// Interfaces shows all interfaces visible in your machine
func AllInterfaces() []string {
	var interfaceSlice []string
	names, err := net.Interfaces()
	if err != nil {
		log.Println(err)
	}

	for _, nm := range names {
		interfaceSlice = append(interfaceSlice, nm.Name)
	}
	return interfaceSlice
}

// ChooseInterface chooses which interface to watch
func ChooseInterface(input []string) string {
	fmt.Println("Enter the number of the interface you would like to watch")
	for i, names := range input {
		fmt.Printf("%d for %s \n", i, names)
	}

	var picked int
	fmt.Scanf("%d", &picked)
	choosen := input[picked]
	return choosen
}

// WatchAll watches all interfaces
func WatchAll(input string) {
	fmt.Println("Check all")
}

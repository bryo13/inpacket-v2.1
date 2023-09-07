package nic

import (
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
func ChooseInterface(input string) string {
	var choosen string
	return choosen
}

// WatchAll watches all interfaces
func WatchAll(input string) {}

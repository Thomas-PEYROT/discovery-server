package services

import (
	"discovery-server/globals"
	"math/rand"
)

// Get a random available port between PortRangeMin and PortRangeMax
func GetNewPort() uint32 {

	// Stocker les ports utilisés
	usedPorts := make(map[uint32]bool)
	for _, instances := range globals.RegisteredMicroservices {
		for _, instance := range instances {
			usedPorts[instance.Port] = true
		}
	}

	// Essayer de trouver un port disponible
	for {
		newPort := uint32(int(globals.PortRangeMin) + rand.Intn(int(globals.PortRangeMax-globals.PortRangeMin+1)))
		if !usedPorts[newPort] {
			return newPort
		}
	}
}

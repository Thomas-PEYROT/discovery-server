package services

import (
	"discovery-server/globals"
	"errors"
	"math/rand"
)

// Get a random available port between PortRangeMin and PortRangeMax
func GetNewPort() (uint32, error) {

	// We store used ports
	usedPorts := make(map[uint32]bool)
	for _, instances := range globals.RegisteredMicroservices {
		for _, instance := range instances {
			usedPorts[instance.Port] = true
		}
	}

	if len(usedPorts) == int(globals.PortRangeMax-globals.PortRangeMin) {
		return 0, errors.New("all ports are already used")
	}

	// We try to find one
	for {
		newPort := uint32(int(globals.PortRangeMin) + rand.Intn(int(globals.PortRangeMax-globals.PortRangeMin)))
		if !usedPorts[newPort] {
			return newPort, nil
		}
	}
}

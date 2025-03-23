package globals

type MicroserviceInstance struct {
	Port uint32
}

var RegisteredMicroservices map[string]map[string]MicroserviceInstance = make(map[string]map[string]MicroserviceInstance)

package globals

type MicroserviceInstance struct {
	Port uint32 `json:"port"`
}

var RegisteredMicroservices map[string]map[string]MicroserviceInstance = make(map[string]map[string]MicroserviceInstance)
var ServerPort uint32
var PortRangeMin uint32
var PortRangeMax uint32

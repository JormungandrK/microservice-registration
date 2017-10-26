package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/JormungandrK/microservice-tools/gateway"

	commonconf "github.com/JormungandrK/microservice-tools/config"
)

// Config holds the microservice full configuration.
type Config struct {
	// Microservice is a gateway.Microservice configuration for self-registration and service config.
	Microservice gateway.MicroserviceConfig `json:"microservice"`

	// Database holds the database configuration
	Database *commonconf.DBConfig `json:"database"`

	// GatewayURL is the URL of the gateway (proxy).
	GatewayURL string `json:"gatewayUrl"`

	// GatewayAdminURL is the administration URL of the API Gateway. Used for purposes of registration of a
	// microservice with the API gateway.
	GatewayAdminURL string `json:"gatewayAdminUrl"`

	// Mail is a map of <service-name>:<service base URL>. For example,
	// "user-microservice": "http://kong.gateway:8001/user"
	Services map[string]string `json:"services"`

	// Services is a map of <property>:<value>. For example,
	// "host": "smtp.example.com"
	Mail map[string]string `json:"mail"`
}

// LoadConfig loads a Config from a configuration JSON file.
func LoadConfig(confFile string) (*Config, error) {
	confBytes, err := ioutil.ReadFile(confFile)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = json.Unmarshal(confBytes, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

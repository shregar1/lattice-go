package config_loader
type LatticeConfig struct {
	ServiceName string `json:"serviceName"`
	Environment string `json:"environment"`
	Port        int    `json:"port"`
}
type LatticeConfigLoader struct{}

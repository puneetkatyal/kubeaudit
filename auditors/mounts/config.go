package mounts

type Config struct {
	SensitivePaths []string `yaml:"denyPathsList"`
	All            bool     `yaml:"allHostPaths"`
}

func (config *Config) GetSensitivePaths() []string {
	if config == nil || len(config.SensitivePaths) == 0 {
		return DefaultSensitivePaths
	}
	return config.SensitivePaths
}

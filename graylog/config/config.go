package config

// Config represents terraform provider's configuration.
type Config struct {
	Endpoint     string
	AuthName     string
	AuthPassword string
	XRequestedBy string
	APIVersion   string
}

func (c *Config) LoadAndValidate() error {
	return nil
}

package config

// Config represents terraform provider's configuration.
type Config struct {
	Endpoint     string
	AuthName     string
	AuthPassword string
	XRequestedBy string
	APIVersion   string
}

func (Config) LoadAndValidate() error {
	return nil
}

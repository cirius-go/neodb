package config

// Config represents the configuration for the service connection
// infrastructure.
type Config struct {
	Debug bool `env:"DEBUG"`
}

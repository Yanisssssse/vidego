package server

type Config struct {
	Host         string
	Port         string
	AuthRequired bool
}

// Suitable for local developement
func DefaultConfig() *Config {
	return &Config{
		Host:         "localhost",
		Port:         "6264",
		AuthRequired: false,
	}
}

func NewConfig(host string, port string, requireAuth bool) *Config {
	return &Config{
		Host:         host,
		Port:         port,
		AuthRequired: requireAuth,
	}
}

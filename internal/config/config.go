package config

type Config struct {
	ConfigVars *configVars
	// Database                     database.DbInterface
	// Logger logrus.FieldLogger
}

func New() (*Config, error) {
	config := Config{}

	// Initialize logger.

	// Connect to db

	return &config, nil
}

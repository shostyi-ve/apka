package config

type Config struct {
	HTTPServerPort string `env:"HTTP_SERVER_PORT"`
}

func NewConfig() (*Config, error) {
	config := &Config{}

	//	if err := envdecode.Decode(config); err != nil {
	//		return nil, errors.WithStack(err)
	//	}

	return config, nil
}
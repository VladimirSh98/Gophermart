package config

import "github.com/caarlos0/env/v6"

func (conf *Config) parseEnv() error {
	err := env.Parse(conf)
	if err != nil {
		return err
	}
	return nil
}

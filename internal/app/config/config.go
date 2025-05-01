package config

func (conf *Config) Load() error {
	var err error

	baseConf := &baseConfig{}
	err = baseConf.parse()
	if err != nil {
		return err
	}

	flagConf := &Config{}
	flagConf.parseFlag(baseConf)

	envConf := &Config{}
	err = envConf.parseEnv()
	if err != nil {
		return err
	}

	conf.result(flagConf, envConf)

	return nil
}

func (conf *Config) result(flagConf *Config, envConf *Config) {
	if envConf.RunAddress != "" {
		conf.RunAddress = envConf.RunAddress
	} else {
		conf.RunAddress = flagConf.RunAddress
	}

	if envConf.AccrualSystemAddress != "" {
		conf.AccrualSystemAddress = envConf.AccrualSystemAddress
	} else {
		conf.AccrualSystemAddress = flagConf.AccrualSystemAddress
	}

	if envConf.DatabaseURI != "" {
		conf.DatabaseURI = envConf.DatabaseURI
	} else {
		conf.DatabaseURI = flagConf.DatabaseURI
	}
}

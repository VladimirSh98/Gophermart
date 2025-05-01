package config

func (conf *Config) Load() error {
	var err error

	baseConf := &baseConfig{}
	err = baseConf.parse()
	if err != nil {
		return err
	}

	flagConf := &flagConfig{}
	flagConf.parse(baseConf)

	envConf := &envConfig{}
	err = envConf.parse()
	if err != nil {
		return err
	}

	conf.result(flagConf, envConf)

	return nil
}

func (conf *Config) result(flagConf *flagConfig, envConf *envConfig) {
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

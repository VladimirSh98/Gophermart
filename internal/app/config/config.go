package config

import "github.com/VladimirSh98/Gophermart.git/internal/app/models"

var Conf models.Config

func Load(conf *models.Config) error {
	var err error

	baseConf := &baseConfig{}
	err = baseConf.parse()
	if err != nil {
		return err
	}

	flagConf := &models.Config{}
	parseFlag(flagConf, baseConf)

	envConf := &models.Config{}
	err = parseEnv(envConf)
	if err != nil {
		return err
	}

	result(conf, flagConf, envConf, baseConf)

	return nil
}

func result(conf *models.Config, flagConf *models.Config, envConf *models.Config, baseConf *baseConfig) {
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

	if envConf.SecretKey != "" {
		conf.SecretKey = envConf.SecretKey
	} else {
		conf.SecretKey = baseConf.SecretKey
	}

	conf.MigrationsDir = baseConf.MigrationsDir
	conf.TokenExp = baseConf.TokenExp
}

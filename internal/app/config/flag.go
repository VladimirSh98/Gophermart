package config

import "flag"

func (conf *flagConfig) parse(baseConf *baseConfig) {
	flag.StringVar(&conf.AccrualSystemAddress, "r", baseConf.AccrualSystemAddress, "Accrual System Address")
	flag.StringVar(&conf.RunAddress, "a", baseConf.RunAddress, "Run address")
	flag.StringVar(&conf.DatabaseURI, "d", baseConf.DatabaseURI, "DB file path")
	flag.Parse()
}

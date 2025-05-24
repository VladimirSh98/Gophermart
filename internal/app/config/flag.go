package config

import (
	"flag"
	"github.com/VladimirSh98/Gophermart.git/internal/app/models"
)

func parseFlag(conf *models.Config, baseConf *baseConfig) {
	flag.StringVar(&conf.AccrualSystemAddress, "r", baseConf.AccrualSystemAddress, "Accrual System Address")
	flag.StringVar(&conf.RunAddress, "a", baseConf.RunAddress, "Run address")
	flag.StringVar(&conf.DatabaseURI, "d", baseConf.DatabaseURI, "DB file path")
	flag.Parse()
}

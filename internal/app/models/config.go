package models

type Config struct {
	AccrualSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	RunAddress           string `env:"RUN_ADDRESS"`
	DatabaseURI          string `env:"DATABASE_URI"`
	SecretKey            string `env:"SECRET_KEY"`
	MigrationsDir        string
	TokenExp             int
}

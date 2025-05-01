package config

type envConfig struct {
	AccrualSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	RunAddress           string `env:"RUN_ADDRESS"`
	DatabaseURI          string `env:"DATABASE_URI"`
}

type baseConfig struct {
	RunAddress           string `yaml:"run_address"`
	AccrualSystemAddress string `yaml:"accrual_system_address"`
	DatabaseURI          string `yaml:"database_uri"`
	MigrationsDir        string `yaml:"migrations_dir"`
}

type flagConfig struct {
	AccrualSystemAddress string
	RunAddress           string
	DatabaseURI          string
}

type Config struct {
	AccrualSystemAddress string
	RunAddress           string
	DatabaseURI          string
}

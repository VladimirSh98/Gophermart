package config

type baseConfig struct {
	RunAddress           string `yaml:"run_address"`
	AccrualSystemAddress string `yaml:"accrual_system_address"`
	DatabaseURI          string `yaml:"database_uri"`
	SecretKey            string `env:"secret_key"`
	MigrationsDir        string `yaml:"migrations_dir"`
	TokenExp             int    `yaml:"token_exp"`
}

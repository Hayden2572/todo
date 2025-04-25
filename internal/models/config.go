package models

type Config struct {
	Port     string   `json:"port" env-required:"true"`
	Address  string   `json:"address" env-required:"true" env-default:"localhost"`
	DataBase DBConfig `json:"database" env-required:"true"`
}

type DBConfig struct {
	DBPort   string `json:"db_port" env-required:"true"`
	Username string `json:"username" env-required:"true"`
	Password string `json:"password" env-required:"true"`
	HostName string `json:"host_name" env-required:"true"`
	DBName   string `json:"db_name" env-required:"true"`
	Sslmode  string `json:"sslmode" env-default:"disabled" env-required:"true"`
}

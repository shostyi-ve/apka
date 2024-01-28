package db

import "fmt"

type DBConfig struct {
	Port     string `env:"DB_PORT"`
	Host     string `env:"DB_HOST"`
	Name     string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASS"`
}

func (db *DBConfig) GetConnectString() string {
	info := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db.Host,
		db.Port,
		db.User,
		db.Password,
		db.Name,
	)

	return info
}

func NewDBConf() *DBConfig {
	return &DBConfig{}
}

package configs

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type DatabaseCredential struct {
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Hostname string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	Database string `env:"DB_NAME"`
	Options  string
}

func SetupDatabase() (connectionString string) {
	var db = DatabaseCredential{
		Options: "timeout=90s&collation=utf8mb4_unicode_ci&parseTime=true",
	}

	env.Parse(db)
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", db.Username, db.Password, db.Hostname, db.Port, db.Database, db.Options)
}

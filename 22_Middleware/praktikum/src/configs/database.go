package configs

import (
	"fmt"
)

type DatabaseCredential struct {
	Username string
	Password string
	Hostname string
	Port     string
	Database string
	Options  string
}

func SetupDatabase() (connectionString string) {
	var serverConfig = GetServerConfig()
	var db = DatabaseCredential{
		Username: serverConfig.DatabaseUsername,
		Password: serverConfig.DatabasePassword,
		Hostname: serverConfig.DatabaseHostname,
		Port:     serverConfig.DatabasePort,
		Database: serverConfig.DatabaseName,
		Options:  "timeout=90s&collation=utf8mb4_unicode_ci&parseTime=true",
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", db.Username, db.Password, db.Hostname, db.Port, db.Database, db.Options)
}

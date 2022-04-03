package configs

import "fmt"

type DatabaseCredential struct {
	Username string
	Password string
	Hostname string
	Port     int
	Database string
	Options  string
}

func SetupDatabase() (connectionString string) {
	var connectionData = DatabaseCredential{
		Username: "root",
		Password: "admin",
		Hostname: "127.0.0.1",
		Port:     3306,
		Database: "km2golangc",
		Options:  "timeout=90s&collation=utf8mb4_unicode_ci&parseTime=true",
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", connectionData.Username, connectionData.Password, connectionData.Hostname, connectionData.Port, connectionData.Database, connectionData.Options)
}

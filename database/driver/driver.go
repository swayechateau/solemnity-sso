package driver

import (
	"database/sql"
	"os"
)

func Connect(dbType string) (*sql.DB, error) {
	switch dbType {
	case "mysql":
		return SqlConnect()
	}
	return SqlConnect()
}

func GetDbUser() string {
	if user := os.Getenv("DB_USER"); user != "" {
		return user
	}
	return "root"
}

func GetDbPassword() string {
	if password := os.Getenv("DB_PASSWORD"); password != "" {
		return password
	}
	return "password"
}

func GetDbHost() string {
	if host := os.Getenv("DB_HOST"); host != "" {
		return host
	}
	return "localhost"
}

func GetDbPort() string {
	if port := os.Getenv("DB_PORT"); port != "" {
		return port
	}
	return "3306"
}

func GetDbName() string {
	if name := os.Getenv("DB_NAME"); name != "" {
		return name
	}
	return "oauth_db"
}

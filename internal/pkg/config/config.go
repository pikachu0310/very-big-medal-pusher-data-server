package config

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func AppAddr() string {
	return getEnv("APP_ADDR", ":8080")
}

func MySQL() *mysql.Config {
	c := mysql.NewConfig()

	_, isDeployedOnNeoShowcase := os.LookupEnv("NS_MARIADB_PORT")
	if isDeployedOnNeoShowcase {
		c.User = getEnv("NS_MARIADB_USER", "root")
		c.Passwd = getEnv("NS_MARIADB_PASSWORD", "pass")
		c.Net = getEnv("DB_NET", "tcp")
		c.Addr = fmt.Sprintf(
			"%s:%s",
			getEnv("NS_MARIADB_HOSTNAME", "localhost"),
			getEnv("NS_MARIADB_PORT", "3306"),
		)
		c.DBName = getEnv("NS_MARIADB_DATABASE", "app")
	} else {
		c.User = getEnv("DB_USER", "root")
		c.Passwd = getEnv("DB_PASSWORD", "pass")
		c.Net = getEnv("DB_NET", "tcp")
		c.Addr = fmt.Sprintf(
			"%s:%s",
			getEnv("DB_HOST", "localhost"),
			getEnv("DB_PORT", "3306"),
		)
		c.DBName = getEnv("DB_NAME", "app")
	}
	c.Collation = "utf8mb4_general_ci"
	c.AllowNativePasswords = true
	c.ParseTime = true

	return c
}

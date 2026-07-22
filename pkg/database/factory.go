package database

import "os"

func CreateDriver(dbType string) IDatabaseDriver {
	if dbType == "" {
		dbType = os.Getenv("DATABASE_TYPE")
	}
	if dbType == "postgres" || dbType == "postgresql" {
		return NewPostgresDriver("postgres://localhost:5432/lattice")
	}
	return NewSqliteDriver(":memory:")
}

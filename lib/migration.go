package lib

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migrations model ...
type Migration struct {
	logger Logger
	env Env
}

// NewMigrations creates new migration instance ...
func NewMigration(
	logger Logger,
	env Env,
) Migration{
	return Migration{
		logger: logger,
		env: env,
	}
}

func (m Migration) Migrate(){

	m.logger.Info("Migrating schemas...........")

	username := m.env.DBUsername
	password := m.env.DBPassword
	host := m.env.DBHost
	port := m.env.DBPort
	dbname := m.env.DBName
	environment := m.env.Environment

	dsn := fmt.Sprintf("%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbname)

	if environment != "local" {
		dsn = m.env.DatabaseURL
	}

	migrations, err := migrate.New("file://migration", "postgres://"+dsn)
	if err != nil {
		m.logger.Error("Error [migration File init] ::", err.Error())
		panic(err)
	}

	m.logger.Info("---- Running Migration -----", dsn)
	err = migrations.Steps(10)
	if err != nil {
		m.logger.Error("Error in migration: ", err.Error())
	}

}
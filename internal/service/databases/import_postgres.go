package databases

import (
	"context"
	"fmt"
	"strconv"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
)

type ImportPostgresParams struct {
	Host     string
	Port     string
	User     string
	Password string
	SSLMode  string
	Version  string
}

func (s *Service) ImportPostgresDatabases(
	ctx context.Context, params ImportPostgresParams,
) (int, error) {
	if params.SSLMode == "" {
		params.SSLMode = "prefer"
	}
	if params.Port == "" {
		params.Port = "5432"
	}

	connString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/postgres?sslmode=%s",
		params.User, params.Password, params.Host, params.Port, params.SSLMode,
	)

	pgVersion, err := s.ints.PGClient.ParseVersion(params.Version)
	if err != nil {
		return 0, fmt.Errorf("invalid postgres version: %w", err)
	}

	err = s.ints.PGClient.Test(pgVersion, connString)
	if err != nil {
		return 0, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	databases, err := s.ints.PGClient.ListDatabases(pgVersion, connString)
	if err != nil {
		return 0, fmt.Errorf("failed to list databases: %w", err)
	}

	importedCount := 0
	for _, dbName := range databases {
		if isSystemDatabase(dbName) {
			continue
		}

		dbConnString := fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
			params.User, params.Password, params.Host, params.Port, dbName, params.SSLMode,
		)

		_, err := s.CreateDatabase(ctx, dbgen.DatabasesServiceCreateDatabaseParams{
			Name:             dbName,
			PgVersion:        params.Version,
			ConnectionString: dbConnString,
		})
		if err != nil {
			continue
		}
		importedCount++
	}

	return importedCount, nil
}

func isSystemDatabase(name string) bool {
	systemDatabases := []string{"postgres", "template0", "template1"}
	for _, sysDB := range systemDatabases {
		if name == sysDB {
			return true
		}
	}
	return false
}

func parsePort(port string) int32 {
	p, err := strconv.ParseInt(port, 10, 32)
	if err != nil {
		return 5432
	}
	return int32(p)
}

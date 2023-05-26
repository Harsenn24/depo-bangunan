package config

import (
	"database/sql"
	"depobangunan/app/environment"
	"depobangunan/app/migration"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func GetDBInstance() (*sql.DB, error) {

	environment.ExportEnv()

	pg_user := os.Getenv("PG_USER")
	pg_password := os.Getenv("PG_PASSWORD")
	pg_db := os.Getenv("PG_DB")
	pg_port := os.Getenv("PG_PORT")
	pg_host := os.Getenv("PG_HOST")

	link_pg1 := "postgres://" + pg_user + ":" + pg_password + "@" + pg_host + ":" + pg_port + "/" + pg_db + "?sslmode=disable"

	db, err := sql.Open("postgres", link_pg1 )
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("connected to postgresql database")

	db.Exec(migration.CreateTableCustomer())
	db.Exec(migration.CreateTableOrder())

	return db, nil
}

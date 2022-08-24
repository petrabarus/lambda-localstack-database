package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func HandleRequest(ctx context.Context) (string, error) {
	const dsn = "postgresql://postgres:postgres@postgres:5432/postgres?sslmode=disable"

	err := migrateUp(dsn)
	if err != nil {
		return "", fmt.Errorf("migration failed: %v", err)
	}

	user_id, err := insert(dsn)

	if err != nil {
		return "", fmt.Errorf("insert failed: %v", err)
	}

	username, err := getUsernameById(dsn, user_id)

	if err != nil {
		return "", fmt.Errorf("get username failed: %v", err)
	}

	return fmt.Sprintf("id: %d, username: %s", user_id, username), nil
}

func migrateUp(dsn string) error {
	path := "file://./migrations"
	migration, err := migrate.New(path, dsn)

	if err != nil {
		return err
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func insert(dsn string) (int, error) {
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return 0, err
	}
	defer db.Close()

	sqlStatement := `
INSERT INTO users (username)
VALUES ( md5(random()::text) )
RETURNING user_id
`
	userId := 0
	err = db.QueryRow(sqlStatement).Scan(&userId)

	if err != nil {
		return 0, err
	}

	return userId, nil
}

func getUsernameById(dsn string, user_id int) (string, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return "", err
	}
	defer db.Close()

	sqlStatement := `
SELECT username FROM users WHERE user_id = $1
`
	var username string
	err = db.QueryRow(sqlStatement, user_id).Scan(&username)

	if err != nil {
		return "", err
	}
	return username, nil
}

func main() {
	lambda.Start(HandleRequest)
}

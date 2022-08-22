package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context) (string, error) {
	const dsn = "postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return "", err
	}
	rows, err := db.Query("SELECT 1 + 4")
	if err != nil {
		return "", err
	}
	rows.Next()
	var result int
	rows.Scan(&result)

	defer db.Close()
	return fmt.Sprintf("result: %d", result), nil
}

func main() {
	lambda.Start(HandleRequest)
}

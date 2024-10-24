package shortener

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func (s *ShortenerUseCase) insert(ctx context.Context, db *sql.DB, original, shortened string) error {
	_, err := db.ExecContext(ctx, "INSERT INTO URL (original_url, shortened_url, created_at) VALUES ($1,$2,$3)",
		original, shortened, time.Now().UTC())
	if err != nil {
		return fmt.Errorf("insert: %w", err)
	}
	return nil
}

// get retrieves the original URL corresponding to the given shortened URL from the database.
// It executes a SQL query to select the original URL where the shortened URL matches the provided query parameter.
// If the query execution or row scanning encounters an error, it returns an empty string and the error.
//
// Parameters:
// - ctx: The context for managing request deadlines and cancellation signals.
// - db: The database connection to execute the query.
// - query: The shortened URL to search for in the database.
//
// Returns:
// - A string containing the original URL if found.
// - An error if any issue occurs during query execution or row scanning.
func (s *ShortenerUseCase) get(ctx context.Context, db *sql.DB, query string) (string, error) {
	rows, err := db.QueryContext(ctx, "SELECT shortened_url FROM URL WHERE original_url = $1", query)
	if err != nil {
		return "", fmt.Errorf("ShortenerUseCase.get.QueryContext: %w", err)
	}
	defer func() {
		_ = rows.Close()
	}()

	if rows.Err() != nil {
		return "", fmt.Errorf("ShortenerUseCase.get.Err: %w", rows.Err())
	}

	var original string
	if rows.Next() {
		err = rows.Scan(&original)
	}
	if err != nil {
		return original, fmt.Errorf("ShortenerUseCase.get.Next: %w", err)
	}

	return original, nil
}

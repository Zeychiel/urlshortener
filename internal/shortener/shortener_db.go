package shortener

import (
	"context"
	"database/sql"
	"time"
)

func (s *ShortenerUseCase) Insert(ctx context.Context, db *sql.DB, original, shortened string) error {
	_, err := db.ExecContext(ctx, "INSERT INTO URL (original_url, shortened_url, created_at) VALUES ($1,$2,$3)", original, shortened, time.Now().UTC())
	return err
}

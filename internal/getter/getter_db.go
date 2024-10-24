package getter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func (s *GetterUseCase) get(ctx context.Context, db *sql.DB, shortened string) (string, error) {
	row := db.QueryRowContext(ctx, "SELECT original_url FROM URL WHERE shortened_url = $1", shortened)

	if row.Err() != nil {
		return "", fmt.Errorf("Get.Err: %w", row.Err())
	}

	var original string
	err := row.Scan(&original)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrNotFound
		}
		return "", fmt.Errorf("Get.Scan: %w", err)
	}

	return original, nil
}

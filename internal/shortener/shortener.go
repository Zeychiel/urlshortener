package shortener

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

var ErrWrongInput = errors.New("wrong input")

// ShortenerUseCase represents the use case for URL shortening
type ShortenerUseCase struct {
	db *sql.DB
}

// NewShortenerUseCase creates a new instance of ShortenerUseCase
func NewShortenerUseCase(db *sql.DB) *ShortenerUseCase {
	return &ShortenerUseCase{
		db: db,
	}
}

// ServeHTTP implements http.Handler.
func (s *ShortenerUseCase) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	input := req.URL.Query().Get("input")
	if input == "" {
		http.Error(w, "Missing input parameter", http.StatusBadRequest)
		return
	}

	err := s.Do(req.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Input registered successfully")
}

// Do processes the input URL, checks its validity, and inserts it into the database if valid
func (s *ShortenerUseCase) Do(ctx context.Context, input string) error {
	if check(input) {
		return s.Insert(ctx, s.db, input, input)
	}
	return ErrWrongInput
}

// check validates the input URL
func check(input string) bool {
	_, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}
	return true
}

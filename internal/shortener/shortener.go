// Package shortener contains the implementation of the ShortenerUseCase struct, which is responsible
// for handling the URL shortening logic.
package shortener

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/google/uuid"

	httpresponse "urlshortener/internal/http"
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
	var p struct {
		Input string `json:"url"`
	}

	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		httpresponse.ResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if p.Input == "" {
		httpresponse.ResponseError(w, "Missing input parameter", http.StatusBadRequest)
		return
	}

	shortened, err := s.Do(req.Context(), p.Input)
	if err != nil {
		httpresponse.ResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpresponse.ResponseOk(w, shortened)
}

// Do processes the input URL, checks its validity, and inserts it into the database if valid
func (s *ShortenerUseCase) Do(ctx context.Context, input string) (string, error) {
	if !check(input) {
		return "", ErrWrongInput
	}

	// Check if already exists in the database
	shortened, err := s.get(ctx, s.db, input)
	if err != nil {
		return "", err
	}
	if shortened != "" {
		return shortened, nil
	}
	// Apply the shortener algorithm
	shortened = uuid.New().String()

	return shortened, s.insert(ctx, s.db, input, shortened)
}

// check validates the input URL
// TODO: sanitize if need be
func check(input string) bool {
	_, err := url.ParseRequestURI(input)
	return err == nil
}

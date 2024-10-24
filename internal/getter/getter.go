// Package getter contains the implementation of the GetterUseCase struct, which is responsible for handling
// the URL shortening process.
package getter

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	httpresponse "urlshortener/internal/http"

	"github.com/gorilla/mux"
)

var (
	ErrWrongInput = errors.New("GetterUseCase wrong input")
)

// GetterUseCase represents the use case for URL shortening
type GetterUseCase struct {
	db *sql.DB
}

// NewGetterUseCase creates a new instance of GetterUseCase
func NewGetterUseCase(db *sql.DB) *GetterUseCase {
	return &GetterUseCase{
		db: db,
	}
}

// ServeHTTP implements http.Handler.
func (s *GetterUseCase) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	input := mux.Vars(req)["shortened_url"]
	if input == "" {
		httpresponse.ResponseError(w, "Missing input parameter", http.StatusBadRequest)
		return
	}

	res, err := s.Do(req.Context(), input)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			httpresponse.ResponseError(w, err.Error(), http.StatusNotFound)
			return
		}
		httpresponse.ResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpresponse.ResponseOk(w, res)
}

// Do processes the input URL, checks its validity, and inserts it into the database if valid
func (s *GetterUseCase) Do(ctx context.Context, input string) (string, error) {
	if isValid(input) {
		return s.get(ctx, s.db, input)
	}
	return "", ErrWrongInput
}

// isValid validates the input URL
func isValid(input string) bool {
	// TODO: implement business validations
	return true
}

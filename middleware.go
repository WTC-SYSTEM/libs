package apperror

import (
	"encoding/json"
	"errors"
	"net/http"
)

type appHandler func(http.ResponseWriter, *http.Request) *AppError

func Middleware(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		err := h(w, r)
		if err != nil {
			if errors.Is(err, ErrNotFound) {
				w.WriteHeader(http.StatusNotFound)
				_ = encoder.Encode(err)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			_ = encoder.Encode(err)
			return
		}
	}
}

package api

import (
	"errors"
	"log"
	"net/http"
)

// HandleError is an middleware for a correct errors rendering.
func HandleError(handler func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			var apiError Error
			if errors.As(err, &apiError) {
				http.Error(w, apiError.Error(), apiError.Code)
			} else {
				log.Printf("err: %v", err)
				http.Error(w, "", http.StatusInternalServerError)
			}
		}
	}
}

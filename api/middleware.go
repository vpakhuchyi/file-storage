package api

import (
	"errors"
	"log"
	"net/http"
)

// WrapRouterError is an middleware for a correct errors rendering.
func WrapRouterError(router func(w http.ResponseWriter, r *http.Request) func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := router(w, r)(w, r); err != nil {
			handleErr(err, w)
		}
	}
}

// WrapHandlerError is an middleware for a correct errors rendering.
func WrapHandlerError(handler func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			handleErr(err, w)
		}
	}
}

func handleErr(err error, w http.ResponseWriter) {
	var apiError Error
	if errors.As(err, &apiError) {
		http.Error(w, apiError.Error(), apiError.Code)
	} else {
		log.Printf("err: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
}

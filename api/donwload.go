package api

import (
	"errors"
	"fmt"
	"net/http"
	"syscall"
)

// Download is a handler that allows files downloading.
func (c Client) Download(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return ErrInvalidRequestMethod
	}

	// File identifier (file name in our case) is passed through query just to simplify the solution.
	// A more convenient way is to take it from the path.
	fileName := r.URL.Query().Get("name")
	if err := c.Storage.Download(fileName, w); err != nil {
		// It's better to mpa errors between different layers to avoid a needs to check for such errors like syscall.ENOENT.
		// However, current option is much faster to implement >_<.
		if errors.Is(err, syscall.ENOENT) {
			return ErrFileNotFound
		}

		return err
	}

	// Content-Type may be more specific instead of "image/*".
	// Much better solution is to store a content-type of an origin file inside the DB and to use that info here.
	w.Header().Set("Content-Type", "image/*")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", fileName))

	return nil
}

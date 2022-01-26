package api

import (
	"bytes"
	"net/http"
)

// List is a handler that allows to get a list of meta-information of all uploaded files.
func (c Client) List(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return ErrInvalidRequestMethod
	}

	meta, err := c.Storage.List()
	if err != nil {
		return err
	}

	buf := bytes.Buffer{}
	for _, s := range meta {
		buf.WriteString(s)
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(buf.Bytes())

	return nil
}

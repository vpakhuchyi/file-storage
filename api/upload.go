package api

import (
	"bytes"
	"net/http"
)

// Upload is a handler that allows to upload an image.
func (c Client) Upload(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return ErrInvalidRequestMethod
	}

	if err := r.ParseMultipartForm(c.Config.MaxFileSize); err != nil {
		return ErrBadInput
	}

	f, header, err := r.FormFile("file")
	if err != nil {
		return ErrBadInput
	}

	if header.Size > c.Config.MaxFileSize {
		return ErrTooBigFile
	}

	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(f); err != nil {
		return ErrBadInput
	}

	if !isValidContentType(buf.Bytes(), c.Config.ValidContentTypes) {
		return ErrInvalidFileContentType
	}

	if err := c.Storage.Save(header.Filename, buf.Bytes()); err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("File has been uploaded successfully"))

	return nil
}

func isValidContentType(data []byte, types []string) bool {
	ct := http.DetectContentType(data)
	for _, v := range types {
		if v == ct {
			return true
		}
	}

	return false
}

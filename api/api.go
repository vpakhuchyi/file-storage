package api

/*
This package ios responsible for images API definition.

In the current implementation the following operations are supported:
1. Upload an image
2. Download an image by its name
3. List existing files meta-information

*/

/*
Thoughts:
	1. It's much better to not use synchronous calls for files uploading. In case of big file it may take hours to upload it.
 	A lot of issues may happen during that time and file may be not uploaded at all.
	In some cases, it's much better to take a task for file uploading and notify the user with a status like: "pending" etc.
	And send one more message to the user once operation will be finished.

	2. In real world applications we have to store file's meta-information somewhere.
	Some generated ID (UUID?) have to be used instead of filename for its lookup.

	3. Files may be stored in some S3-like blob storage.

	4. Permissions model is not defined here. All users have access to all files. That's not good.
*/

import (
	"io"

	"file-storage/config"
)

// Client is a struct that shall be used for interactions with API.
type Client struct {
	Config  config.API
	Storage Storage
}

// Storage defines and interface between API and any compatible storage.
type Storage interface {
	Save(fileName string, content []byte) error
	Download(fileName string, dest io.Writer) error
	List() ([]string, error)
}

// New is a constructor of API Client.
func New(cfg config.API, storage Storage) Client {
	return Client{
		Config:  cfg,
		Storage: storage,
	}
}

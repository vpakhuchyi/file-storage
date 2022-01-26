package config

/*
This package ios responsible for configuration collection and description.

In the current implementation it's hard-coded inside the constructor.
In real life it' much better to move it to ENV variables and to retrieve it from there.
*/

// Config contains app's configuration.
type Config struct {
	API     API
	Storage Storage
	Server  Server
}

// Server defines configuration that is needed for server.
type Server struct {
	Host string
	Port string
}

// Storage defines configuration that is needed for Storage.
type Storage struct {
	FileDir string
}

// API defines configuration that is needed for API.
type API struct {
	MaxFileSize       int64
	ValidContentTypes []string
}

// New is a constructor of Config.
func New() Config {
	return Config{
		API: API{
			ValidContentTypes: []string{"image/jpeg", "image/jpg", "image/gif", "image/png"},
			MaxFileSize:       10485760, // 10 MB
		},
		Server: Server{
			Host: "localhost",
			Port: "8080",
		},
		Storage: Storage{
			FileDir: "files/",
		},
	}
}

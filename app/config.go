// +build !appengine

package main

type (
	// Config is an example of provider-specific configuration
	Config struct {
		MongoURL          string
		SQLiteStoragePath string
	}
)

var (
	config *Config
)

func init() {
	// this would likely be loaded from flags or a conf file
	config = &Config{
		MongoURL:          "mongodb://localhost/clean",
		SQLiteStoragePath: "/var/tmp",
	}
}

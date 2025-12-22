package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ReleaseVersion = ""
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot load environment file " + err.Error())
	}
}

func Load() {
	loadEnv("RELEASE_VERSION", &ReleaseVersion)

}
func loadEnv(key string, val *string) {
	if *val == "" {
		*val = os.Getenv(key)
	}
}

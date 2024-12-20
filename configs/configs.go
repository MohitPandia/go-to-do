package configs

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func Loadconfigs() {
	_, b, _, _ := runtime.Caller(0)
	rootPath := filepath.Join(filepath.Dir(b), "../")

	err := godotenv.Load(rootPath + "/.env")
	if err != nil {
		// log.Fatal("error loading .env file")
		log.Println("error loading .env file") // this change is made just to make it work on dev and staging as variables will be directly on kubernetes
	}

	loadDBConfig()
}

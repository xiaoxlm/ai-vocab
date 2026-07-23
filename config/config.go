package config

import (
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func init() {
	rootDir := projectRoot()
	if err := godotenv.Load(filepath.Join(rootDir, ".env")); err != nil {
		panic(err)
	}
}

func projectRoot() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to get current file path")
	}

	return filepath.Dir(file)
}

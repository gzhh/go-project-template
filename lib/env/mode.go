package env

import (
	"os"
)

const (
	RunMode = "RUN_MODE"

	DevMode  = "dev"
	TestMode = "test"
	ProdMode = "prod"
)

var modeName = "dev"

func init() {
	mode := os.Getenv(RunMode)
	if mode != "" {
		SetMode(mode)
	}
}

func SetMode(mode string) {
	modeName = mode
}

func Mode() string {
	return modeName
}

func IsDev() bool {
	return modeName == DevMode
}

func IsTest() bool {
	return modeName == TestMode
}

func IsProd() bool {
	return modeName == ProdMode
}

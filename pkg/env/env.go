package env

import (
	"fmt"
	"os"
)

func Check(env string) string {
	if os.Getenv(env) != "" {
		return os.Getenv(env)
	}

	tail := "not set... I hope you know what you're doing! \n\n"
	fmt.Printf(env + " " + tail)
	return ""
}

func SetDefault(env string, url string) string {
	if os.Getenv(env) != "" {
		return os.Getenv(env)
	}

	tail := "not set, using: " + url + " \n\n"
	fmt.Printf(env + " " + tail)
	return url
}

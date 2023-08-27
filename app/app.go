package app

import "os"

func Url(route string) string {
	url := os.Getenv("URL")
	if url == "" {
		url = "http://localhost:" + Port() + "/"
	}

	return url + route
}

func Port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

package server

import "net/http"

func Serve() error {
	http.HandleFunc("/", handleIdenticon)

	// Google Chrome requests for /favicon.ico when accessing a page
	http.HandleFunc("/favicon.ico", handleEmpty)

	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

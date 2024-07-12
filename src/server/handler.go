package server

import (
	"net/http"

	"github.com/pasca-l/identicon-generator/identicon"
)

func handleEmpty(w http.ResponseWriter, r *http.Request) {}

func handleIdenticon(w http.ResponseWriter, r *http.Request) {
	// get rid of leading "/" in url path
	userName := r.URL.Path[1:]

	w.Header().Set("Content-Type", "image/svg+xml")
	identicon.GenerateIdenticon(userName, w)
}

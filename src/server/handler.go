package server

import (
	"net/http"

	"github.com/pasca-l/identicon-svg-generator/identicon"
)

func handleEmpty(w http.ResponseWriter, r *http.Request) {}

func handleIdenticon(w http.ResponseWriter, r *http.Request) {
	// get rid of leading "/" in url path
	userName := r.URL.Path[1:]

	icon, err := identicon.GenerateIdenticon(userName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	identicon.DrawIdenticon(w, icon)
}

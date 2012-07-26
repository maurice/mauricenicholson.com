// Copyright Maurice Nicholson. All rights reserved.

package app

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}
	http.ServeFile(w, r, "templates/index.html")
}

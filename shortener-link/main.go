package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter/"
)

func main() {
	router := httprouter.New()
	router.GET("/:id", handleShortLink)
	http.ListenAndServe(":8080", router)
}

func handleShortLink(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	// Lookup the ID in a database or map and redirect to the original URL
	url, ok := links[id]
	if !ok {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

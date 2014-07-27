package api

import (
	"net/http"
)

func init() {
	http.Handle("/", AppHandler(handler))
}

func handler(w http.ResponseWriter, s *Scope) {
	// do something?
}

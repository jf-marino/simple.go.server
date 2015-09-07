package routes

import "net/http"

type RootHandler struct{}

func (root *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You are now in the root of the app"))
}

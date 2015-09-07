package routes

import "net/http"

type PeopleHandler struct{}

func (root *PeopleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You are now in the people page..."))
}

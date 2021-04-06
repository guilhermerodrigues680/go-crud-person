package controller

import "net/http"

func RootHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("API v1 ok!"))
}

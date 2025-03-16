package main

import "net/http"

func handlerFirst(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{}{})
}

package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondToJson(w,200,struct{}{})
}
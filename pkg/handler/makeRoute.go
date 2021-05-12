package handler

import (
	"crudSystem/pkg/utils"
	"net/http"
)

func MakeRoute(method, url string, fn func(http.ResponseWriter, *http.Request, string)) {
	http.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			fn(w, r, url)
		} else {
			utils.ServerResponse(w, http.StatusNotFound)
			return
		}
	})
}

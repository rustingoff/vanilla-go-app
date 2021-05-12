package utils

import (
	"fmt"
	"net/http"
)

func ServerResponse(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf("%d ", status) + http.StatusText(status)))
}

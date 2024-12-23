package handlers

import (
	"fmt"
	"net/http"
)

// ErrorHandler handles different HTTP errors
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/html")
	switch status {
	case http.StatusBadRequest:
		fmt.Fprintf(w, "<h1>400 - Bad Request</h1><p>%s</p>", message)
	case http.StatusNotFound:
		fmt.Fprintf(w, "<h1>404 - Not Found</h1><p>%s</p>", message)
	case http.StatusInternalServerError:
		fmt.Fprintf(w, "<h1>500 - Internal Server Error</h1><p>%s</p>", message)
	default:
		fmt.Fprintf(w, "<h1>%d - %s</h1><p>%s</p>", status, http.StatusText(status), message)
	}
}

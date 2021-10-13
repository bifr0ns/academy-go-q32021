package error

import (
	"encoding/json"
	"net/http"
)

//FormattedError is used to return an error in json format with the field message.
type FormattedError struct {
	Message string `json:"message"`
}

//NotFoundHandler will return an error in json format for every route not declared.
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(FormattedError{Message: "The page you requested could not be found."})
}

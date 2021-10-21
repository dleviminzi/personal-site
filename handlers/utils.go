package handlers

import (
	"encoding/json"
	"net/http"
)

// SendJSON takes a data model and name and returns the model keyed to the name in json form
func SendJSON(w http.ResponseWriter, status int, data interface{}, name string) error {
	toSend := make(map[string]interface{})
	toSend[name] = data

	jsn, err := json.Marshal(toSend)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsn)

	return nil
}

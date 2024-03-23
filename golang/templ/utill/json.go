package utill

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/helmimuzkr/templ/types"
)

func JsonDecode(r *http.Request, target any) error {
	if r.Body == nil {
		return fmt.Errorf("error during JsonDecode - missing request body")
	}

	if err := json.NewDecoder(r.Body).Decode(target); err != nil {
		return fmt.Errorf("error during JsonDecode - decode json - %s", err)
	}

	return nil
}

func JsonEncodeWriter(w http.ResponseWriter, status int, source any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	response := types.Data{
		Status:  http.StatusText(status),
		Content: source,
	}

	return json.NewEncoder(w).Encode(response)
}

func JsonErrorWriter(w http.ResponseWriter, status int, message string) error {
	return JsonEncodeWriter(w, status, map[string]string{"message": message})
}

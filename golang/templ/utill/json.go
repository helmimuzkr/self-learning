package utill

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/helmimuzkr/templ/types"
)

const (
	defaultErrMsg = "error during %s - %s"
)

func JsonDecode(r *http.Request, target any) error {
	if r.Body == nil {
		err := fmt.Errorf("missing request body")
		slog.Error(defaultErrMsg, "jsonDecode", err)
		return err
	}

	if err := json.NewDecoder(r.Body).Decode(target); err != nil {
		slog.Error(defaultErrMsg, "jsonDecode", err)
		return fmt.Errorf("error during  decode json")
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

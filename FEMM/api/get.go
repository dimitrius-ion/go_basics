package api

import (
	"encoding/json"
	"net/http"
	"github.com/dimitrius-ion/go_basics/femm/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.GetAll())
}
package api

import (
	"encoding/json"
	"net/http"
	"github.com/dimitrius-ion/go_basics/femm/data"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id") // this is the same as the next line
	// id := r.URL.Query()["id"]
	
	if id != "" {
		index, err := strconv.Atoi(id)
		print(index)
		if err == nil && index >= 0 && index < len(data.GetAll()) {
			json.NewEncoder(w).Encode(data.GetAll()[index])
		}else{
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	}else{
		json.NewEncoder(w).Encode(data.GetAll())
	}
}
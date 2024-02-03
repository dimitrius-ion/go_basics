package api

import (
	"encoding/json"
	"net/http"
	"github.com/dimitrius-ion/go_basics/femm/data"
)


func Post (w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var thing data.Things
		err := json.NewDecoder(r.Body).Decode(&thing)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data.AddThing(thing)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Created"))

	}else{
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
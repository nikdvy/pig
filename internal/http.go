package internal

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/rs/zerolog/log"
)

func RunServer() {
	r := mux.NewRouter()
	r.HandleFunc("/internal/healthz", healthcheck)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("resources/"))))
	http.Handle("/", r)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal().Err(err).Msg("Startup failed")
	}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("Triggering healthcheck")
	var response API
	if r.Method != http.MethodGet {
		response = API{
			Code:    405,
			Message: "Method not allowed",
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		response = API{
			Code:    200,
			Message: "healthy",
		}
		w.WriteHeader(http.StatusOK)
	}
	js, err := json.Marshal(response)
	if err != nil {
		log.Debug().Msg(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Server", "PIG Dummy Service")
	w.Write(js)
}

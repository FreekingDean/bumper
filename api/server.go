package api

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"

	"github.com/FreekingDean/bumper/pkg/database"
)

var db *database.Database

// Start is a blocking call to start the API
// server.
func Start() {
	var err error
	db, err = database.Open()
	defer db.Close()
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.Use(stdLogger)
	r.Methods("OPTIONS").Handler(stdLogger(controlHeader(http.HandlerFunc(notFound))))
	api := r.PathPrefix("/api").Subrouter()
	api.Use(controlHeader)
	api.HandleFunc("/searcher", handleSearch)
	api.HandleFunc("/subscriptions", handleSubscribe).Methods("POST")
	api.HandleFunc("/subscriptions", handleGetSubscribtions).Methods("GET")
	http.ListenAndServe("0.0.0.0:3001", r)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte("Not found"))
}

func stdLogger(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
}

func controlHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := w.Header()
		headers.Set("Access-Control-Allow-Origin", "*")
		headers.Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
		} else {
			h.ServeHTTP(w, r)
		}
	})
}

func respondErr(err error, w http.ResponseWriter) {
	fmt.Printf("ERROR: %+v\n", err)
	w.WriteHeader(500)
	w.Write([]byte(`{"error":"` + err.Error() + `"}`))
}

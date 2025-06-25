package router

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vladimir/note-api/handlers"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/notes", handlers.GetAllNotes(db)).Methods("GET")
	r.HandleFunc("/notes", handlers.CreateNote(db)).Methods("POST")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Note API is running"))
	})

	return r
}

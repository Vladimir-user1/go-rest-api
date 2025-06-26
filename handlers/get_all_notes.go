package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vladimir/note-api/db"
)

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := db.GetAllNotes()
	if err != nil {
		http.Error(w, "Ошибка при получении заметок", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(notes); err != nil {
		http.Error(w, "Ошибка при кодировании ответа", http.StatusInternalServerError)
	}
}

package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/vladimir/note-api/models"
)

func GetAllNotes(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, title, content, created_at FROM notes ORDER BY created_at DESC")
		if err != nil {
			http.Error(w, "Failed to query notes: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var notes []models.Note

		for rows.Next() {
			var note models.Note
			err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt)
			if err != nil {
				http.Error(w, "Error reading notes", http.StatusInternalServerError)
				return
			}
			notes = append(notes, note)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notes)
	}
}

package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/vladimir/note-api/models"
)

func CreateNote(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Начался обработчик POST /notes")

		var input struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		}

		// Декодируем тело запроса
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Неверный JSON", http.StatusBadRequest)
			return
		}

		// Показываем, что получили от клиента
		log.Println("Получен ввод:", input.Title, input.Content)

		// Выполняем вставку
		var note models.Note
		err := db.QueryRow(
			`INSERT INTO notes (title, content) VALUES ($1, $2)
			RETURNING id, title, content, created_at`,
			input.Title, input.Content,
		).Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt)

		if err != nil {
			http.Error(w, "Ошибка вставки: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Показываем, что вернем клиенту
		log.Println("Создана заметка:", note)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(note)
	}

}

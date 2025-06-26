// db/db.go
package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vladimir/note-api/models"
)

var DB *sql.DB

func Connect() *sql.DB {
	// Загружаем переменные из .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env:", err)
	}

	// Получаем переменные
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("БД недоступна:", err)
	}

	log.Println("Успешное подключение в PostgreSQL")
	return DB
}

func GetAllNotes() ([]*models.Note, error) {
	rows, err := DB.Query("SELECT id, title, content, created_at FROM notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []*models.Note
	for rows.Next() {
		var note models.Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt); err != nil {
			return nil, err
		}
		notes = append(notes, &note)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}

// Получение заметки по ID

func GetNoteByID(id int) (*models.Note, error) {
	var note models.Note
	row := DB.QueryRow("SELECT id, title, content, created_at FROM notes WHERE id = $1", id)
	err := row.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &note, nil
}

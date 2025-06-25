package main

import (
	"log"
	"net/http"

	"github.com/vladimir/note-api/db"
	"github.com/vladimir/note-api/router"
)

func main() {
	// Подключение к БД и сохранение объекта подключения
	database := db.Connect()

	// Настройка маршрутов с передачей подключения к БД
	r := router.SetupRoutes(database)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

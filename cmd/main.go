package main

import (
	"cinema/internal/controllers"
	"cinema/internal/repositories"
	"cinema/internal/servises"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Cinema API running")

	// Ініціалізація логгера
	logger := logrus.New()
	logger.Out = os.Stdout

	// Ініціалізація репозиторіїв
	repo := repositories.NewMovieRepository()

	// Ініціалізація служб
	service := services.NewMovieService(repo)

	// Ініціалізація контролера
	controller := controllers.NewMovieController(service)

	// Ініціалізація роутера
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Роут для отримання списку доступних фільмів
	r.Get("/movies", controller.ListMovies)

	// Роут для замовлення квитків
	r.Post("/order", controller.OrderMovie)

	// Роут для отримання списку замовлених квитків
	r.Get("/orders", controller.ListOrders)

	// Запуск сервера на динамічному порту
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logger.Infof("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		logger.Fatal("Failed to start server:", err)
	}
}


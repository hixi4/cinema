package main

import (
	"cinema/internal/controller"
	"cinema/internal/repository"
	"cinema/internal/service"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	// Ініціалізація логгера
	logger := logrus.New()
	logger.Out = os.Stdout

	// Ініціалізація репозиторіїв
	repo := repository.NewMovieRepository()

	// Ініціалізація служб
	emailService := &service.EmailService{} // ваш реальний сервіс
	movieService := service.NewMovieService(repo, emailService)

	// Ініціалізація контролера
	controller := controller.NewMovieController(movieService)

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

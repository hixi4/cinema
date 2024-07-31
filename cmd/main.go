package main

import (
	"cinema/internal/controller"
	"cinema/internal/repository"
	"cinema/internal/service"
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
	repo := repository.NewMovieRepository()

	// Ініціалізація служб
	emailService := service.NewEmailService()
	svc := service.NewMovieService(repo, emailService)

	// Ініціалізація контролера
	ctrl := controller.NewMovieController(svc)

	// Ініціалізація роутера
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Роут для отримання списку доступних фільмів
	r.Get("/movies", ctrl.ListMovies)

	// Роут для замовлення квитків
	r.Post("/order", ctrl.OrderMovie)

	// Роут для отримання списку замовлених квитків
	r.Get("/orders", ctrl.ListOrders)

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

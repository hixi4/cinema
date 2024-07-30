package services

import (
	"cinema/internal/models"
	"cinema/internal/repositories"
	"errors"
	"fmt"
	"log"
	"time"
)

// MovieService структура для служби фільмів
type MovieService struct {
	repo *repositories.MovieRepository
}

// OrderRequest структура для запиту замовлення
type OrderRequest struct {
	MovieTitle string `json:"movie_title"`
}

// NewMovieService створює новий екземпляр MovieService
func NewMovieService(repo *repositories.MovieRepository) *MovieService {
	return &MovieService{repo: repo}
}

// GetAvailableMovies повертає список доступних фільмів
func (s *MovieService) GetAvailableMovies() []models.Movie {
	return s.repo.GetAvailableMovies()
}

// PlaceOrder створює нове замовлення
func (s *MovieService) PlaceOrder(req OrderRequest) (string, error) {
	if req.MovieTitle == "" {
		return "", errors.New("movie title is required")
	}

	orderID := generateOrderID()
	order := models.Order{
		ID:         orderID,
		MovieTitle: req.MovieTitle,
		Status:     "Ordered",
		OrderedAt:  time.Now(),
	}
	s.repo.PlaceOrder(order)
	logOrderEmail(orderID)
	return orderID, nil
}

// GetOrders повертає список замовлених фільмів
func (s *MovieService) GetOrders() []models.Order {
	return s.repo.GetOrders()
}

// generateOrderID генерує унікальний ідентифікатор замовлення
func generateOrderID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// logOrderEmail логує відправку email після замовлення фільма
func logOrderEmail(orderID string) {
	log.Printf("Sent confirmation email for order ID: %s", orderID)
}


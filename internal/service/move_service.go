package service

import (
	"cinema/internal/model"
	"cinema/internal/repository"
	"errors"
	"fmt"
	"log"
	"time"
)

// MovieServiceInterface визначає інтерфейс для служби фільмів
type MovieServiceInterface interface {
	GetAvailableMovies() []model.Movie
	PlaceOrder(req OrderRequest) (string, error)
	GetOrders() []model.Order
}

// MovieService структура для служби фільмів
type MovieService struct {
	repo         *repository.MovieRepository
	emailService EmailServiceInterface
}

// OrderRequest структура для запиту замовлення
type OrderRequest struct {
	MovieTitle string `json:"movie_title"`
}

// NewMovieService створює новий екземпляр MovieService
func NewMovieService(repo *repository.MovieRepository, emailService EmailServiceInterface) *MovieService {
	return &MovieService{repo: repo, emailService: emailService}
}

// GetAvailableMovies повертає список доступних фільмів
func (s *MovieService) GetAvailableMovies() []model.Movie {
	return s.repo.GetAvailableMovies()
}

// PlaceOrder створює нове замовлення
func (s *MovieService) PlaceOrder(req OrderRequest) (string, error) {
	if req.MovieTitle == "" {
		return "", errors.New("movie title is required")
	}

	orderID := generateOrderID()
	order := model.Order{
		ID:         orderID,
		MovieTitle: req.MovieTitle,
		Status:     "Ordered",
		OrderedAt:  time.Now(),
	}
	s.repo.PlaceOrder(order)
	s.emailService.SendOrderEmail(orderID)
	return orderID, nil
}

// GetOrders повертає список замовлених фільмів
func (s *MovieService) GetOrders() []model.Order {
	return s.repo.GetOrders()
}

// generateOrderID генерує унікальний ідентифікатор замовлення
func generateOrderID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

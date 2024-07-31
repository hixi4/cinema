package repositories

import (
	"cinema/internal/models"
)

// MovieRepository структура для репозиторію фільмів
type MovieRepository struct {
	movies []models.Movie
	orders []models.Order
}

// NewMovieRepository створює новий екземпляр MovieRepository
func NewMovieRepository() *MovieRepository {
	return &MovieRepository{
		movies: []models.Movie{
			{Title: "Movie 1", Description: "Description 1", Price: 10.0, Transport: "Bus"},
			{Title: "Movie 2", Description: "Description 2", Price: 12.0, Transport: "Train"},
		},
		orders: make([]models.Order, 0),
	}
}

// GetAvailableMovies повертає список доступних фільмів
func (r *MovieRepository) GetAvailableMovies() []models.Movie {
	return r.movies
}

// PlaceOrder створює нове замовлення
func (r *MovieRepository) PlaceOrder(order models.Order) {
	r.orders = append(r.orders, order)
}

// GetOrders повертає список замовлених фільмів
func (r *MovieRepository) GetOrders() []models.Order {
	return r.orders
}


package repository

import (
	"cinema/internal/model"
)

// MovieRepository структура для репозиторію фільмів
type MovieRepository struct {
	movies []model.Movie
	orders []model.Order
}

// NewMovieRepository створює новий екземпляр MovieRepository
func NewMovieRepository() *MovieRepository {
	return &MovieRepository{
		movies: []model.Movie{
			{Title: "Movie 1", Description: "Description 1", Price: 10.0, Transport: "Bus"},
			{Title: "Movie 2", Description: "Description 2", Price: 12.0, Transport: "Train"},
		},
		orders: make([]model.Order, 0),
	}
}

// GetAvailableMovies повертає список доступних фільмів
func (r *MovieRepository) GetAvailableMovies() []model.Movie {
	return r.movies
}

// PlaceOrder створює нове замовлення
func (r *MovieRepository) PlaceOrder(order model.Order) {
	r.orders = append(r.orders, order)
}

// GetOrders повертає список замовлених фільмів
func (r *MovieRepository) GetOrders() []model.Order {
	return r.orders
}



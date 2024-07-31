package controllers

import (
	"encoding/json"
	"net/http"

	"cinema/internal/servises"
)

// MovieController структура для контролера фільмів
type MovieController struct {
	service *services.MovieService
}

// NewMovieController створює новий екземпляр MovieController
func NewMovieController(service *services.MovieService) *MovieController {
	return &MovieController{service: service}
}

// ListMovies обробляє запит для отримання списку доступних фільмів
func (c *MovieController) ListMovies(w http.ResponseWriter, r *http.Request) {
	movies := c.service.GetAvailableMovies()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		http.Error(w, "Failed to encode movies", http.StatusInternalServerError)
	}
}

// OrderMovie обробляє запит для замовлення фільму
func (c *MovieController) OrderMovie(w http.ResponseWriter, r *http.Request) {
	var req services.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	orderID, err := c.service.PlaceOrder(req)
	if err != nil {
		http.Error(w, "Failed to place order", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(orderID))
}

// ListOrders обробляє запит для отримання списку замовлених фільмів
func (c *MovieController) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders := c.service.GetOrders()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(orders); err != nil {
		http.Error(w, "Failed to encode orders", http.StatusInternalServerError)
	}
}


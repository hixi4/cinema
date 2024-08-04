package controller

import (
	"cinema/internal/service"
	"encoding/json"
	"net/http"
)

// MovieServiceInterface визначає інтерфейс для служби фільмів
type MovieServiceInterface interface {
	GetAvailableMovies() []model.Movie
	PlaceOrder(req service.OrderRequest) (string, error)
	GetOrders() []model.Order
}

type MovieController struct {
	service MovieServiceInterface
}

func NewMovieController(service MovieServiceInterface) *MovieController {
	return &MovieController{service: service}
}

func (c *MovieController) ListMovies(w http.ResponseWriter, r *http.Request) {
	movies := c.service.GetAvailableMovies()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func (c *MovieController) OrderMovie(w http.ResponseWriter, r *http.Request) {
	var req service.OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	orderID, err := c.service.PlaceOrder(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"order_id": orderID})
}

func (c *MovieController) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders := c.service.GetOrders()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

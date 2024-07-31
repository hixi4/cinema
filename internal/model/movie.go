package model

import "time"

// Movie структура для зберігання інформації про фільм
type Movie struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Transport   string  `json:"transport"`
}

// Order структура для зберігання інформації про замовлення
type Order struct {
	ID         string    `json:"id"`
	MovieTitle string    `json:"movie_title"`
	Status     string    `json:"status"`
	OrderedAt  time.Time `json:"ordered_at"`
}

package common

import "math"

type Resultset struct {
	TotalItems   int         `json:"total_items" db:"total_items"`
	TotalPages   int         `json:"total_pages"`
	ItemsPerPage int         `json:"items_per_page"`
	Page         int         `json:"page"`
	Offset       int         `json:"-"`
	Data         interface{} `json:"data"`
}

func (r *Resultset) Calculate() {
	r.TotalPages = int(math.Ceil(float64(r.TotalItems) / float64(r.ItemsPerPage)))

	r.Offset = (r.Page - 1) * r.ItemsPerPage
}

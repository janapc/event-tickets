package pagination

import "math"

type Pagination struct {
	Page       int `json:"page"`
	Size       int `json:"size"`
	TotalItems int `json:"total"`
	TotalPages int `json:"totalPages"`
}

func NewPagination(page, size, totalItems int) Pagination {
	totalPages := 1
	if size > 0 {
		totalPages = int(math.Ceil(float64(totalItems) / float64(size)))
	}
	return Pagination{
		Page:       page,
		Size:       size,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}
}

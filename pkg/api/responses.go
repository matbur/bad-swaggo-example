package api

type PaginatedResponse[T any] struct {
	PageNumber uint64 `json:"page"`
	PageSize   uint64 `json:"perPage"`
	Total      uint64 `json:"total"`
	TotalPages uint64 `json:"totalPages"`
	Data       []T    `json:"data"`
}

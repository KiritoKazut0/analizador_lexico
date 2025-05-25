package entities


type PaginatedUsersResponse struct {
	Users       []User `json:"users"`
	CurrentPage int    `json:"current_page"`
	TotalPages  int    `json:"total_pages"`
	TotalCount  int64  `json:"total_count"`
	PerPage     int    `json:"per_page"`
}

type BatchCreateRequest struct {
	Users []User `json:"users"`
}
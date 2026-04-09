package response

type Meta struct {
	Page  int    `json:"page,omitempty"`
	Size  int    `json:"size,omitempty"`
	Sort  string `json:"sort,omitempty"`
	Total int    `json:"total,omitempty"`
}

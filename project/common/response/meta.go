package response

type Meta struct {
	Page   int    `json:"page,omitempty" min:"1"`
	Size   int    `json:"size,omitempty" min:"1" max:"100"`
	Sort   string `json:"sort,omitempty"`
	SortBy string `json:"sort_by,omitempty" enum:"asc,desc"`
	Total  int64  `json:"total,omitempty"`
}

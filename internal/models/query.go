// internal/models/query.go
package models

// PaginationParams represents pagination parameters
type PaginationParams struct {
	Page     int `json:"page" query:"page"`
	PageSize int `json:"page_size" query:"page_size"`
}

// GetOffset calculates the offset for database queries
func (p *PaginationParams) GetOffset() int {
	if p.Page < 1 {
		p.Page = 1
	}
	return (p.Page - 1) * p.PageSize
}

// GetLimit returns the limit for database queries
func (p *PaginationParams) GetLimit() int {
	if p.PageSize < 1 {
		p.PageSize = 10
	}
	if p.PageSize > 100 {
		p.PageSize = 100
	}
	return p.PageSize
}

// UserFilterParams represents filters for user queries
type UserFilterParams struct {
	PaginationParams
	Search   string `json:"search,omitempty" query:"search"`
	IsActive *bool  `json:"is_active,omitempty" query:"is_active"`
	RoleID   string `json:"role_id,omitempty" query:"role_id"`
}
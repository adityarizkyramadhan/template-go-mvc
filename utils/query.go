package utils

type Query struct {
	Search  string `form:"search"`
	Page    int    `form:"page"`
	PerPage int    `form:"per_page"`
}

type PaginateData struct {
	Rows        interface{} `json:"rows"`
	Total       int64       `json:"total"`
	TotalPages  int         `json:"total_pages"`
	CurrentPage int         `json:"current_page"`
	PrevPage    int         `json:"prev_page"`
	NextPage    int         `json:"next_page"`
}

func Paginate(data interface{}, total int64, page int, perPage int) PaginateData {
	totalPages := int(total) / perPage
	if int(total)%perPage > 0 {
		totalPages++
	}

	return PaginateData{
		Rows:        data,
		Total:       total,
		TotalPages:  safetyZero(totalPages),
		CurrentPage: safetyZero(page),
		PrevPage:    safetyZero(page - 1),
		NextPage:    safetyZero(page + 1),
	}
}

func safetyZero(value int) int {
	if value <= 0 {
		return 1
	}
	return value
}

func AllowedStatus(role string, status string) bool {
	if role == "admin" {
		return true
	}

	if role == "user" && status == "published" {
		return true
	}

	return false
}

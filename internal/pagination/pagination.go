package pagination

import (
	"net/http"
	"strconv"
)

type Pagination struct {
	Page     int
	PageSize int
}

func (p *Pagination) LimitParams() int {
	switch {
	case p.PageSize <= 10:
		return 10
	case p.PageSize <= 25:
		return 25
	default:
		return 50
	}
}

func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.LimitParams()
}

func (p *Pagination) PageParams(pageStr, pageSizeStr string) error {
	var err error

	if pageStr == "" {
		p.Page = 1
	} else {
		p.Page, err = strconv.Atoi(pageStr)
		if err != nil {
			return err
		}
	}

	if p.PageSize, err = strconv.Atoi(pageSizeStr); err != nil {
		return err
	}
	return nil
}

func (p *Pagination) ParsePaginationParams(r *http.Request) (pagination *Pagination) {
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("limit")

	p.PageParams(pageStr, pageSizeStr)
	p.PageSize = p.LimitParams()

	return pagination
}
package pagination_test

import (
	"dashboard/internal/pagination"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePaginationParams(t *testing.T) {
	req := &http.Request{URL: &url.URL{}}
	p := pagination.Pagination{}
	p.ParsePaginationParams(req)
	assert.Equal(t, 1, p.Page, "Page deve ser 1 por padrão")
	assert.Equal(t, 10, p.PageSize, "PageSize deve ser 10 por padrão")

	req = &http.Request{URL: &url.URL{RawQuery: "page=2&limit=20"}}
	p = pagination.Pagination{}
	p.ParsePaginationParams(req)
	assert.Equal(t, 2, p.Page, "Page deve ser 2")
	assert.Equal(t, 25, p.PageSize, "PageSize deve ser 25")

	req = &http.Request{URL: &url.URL{RawQuery: "page=3&limit=abc"}}
	p = pagination.Pagination{}
	p.ParsePaginationParams(req)
	assert.Equal(t, 3, p.Page, "Page deve ser 1")
	assert.Equal(t, 10, p.PageSize, "PageSize deve ser 10 (valor padrão)")
}
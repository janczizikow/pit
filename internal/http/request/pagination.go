package request

import (
	"fmt"
	"math"
	"strings"

	"github.com/janczizikow/pit/internal/validator"
)

const paginationMaxPageSize = 1000

type paginator struct {
	page        int
	pageSize    int
	sort        []string
	sortSafeMap map[string]bool
}

func NewPaginator(pageSize, page int, sort []string, sortSafeMap map[string]bool) *paginator {
	return &paginator{pageSize: pageSize, page: page, sort: sort, sortSafeMap: sortSafeMap}
}

func (p paginator) Limit() int {
	return p.pageSize
}

func (p paginator) Offset() int {
	return (p.page - 1) * p.pageSize
}

func (p paginator) Sort() string {
	var sort []string
	for _, value := range p.sort {
		if _, ok := p.sortSafeMap[value]; ok {
			if strings.HasPrefix(value, "-") {
				sort = append(sort, fmt.Sprintf("%s %s", strings.TrimPrefix(value, "-"), "DESC"))
			} else {
				sort = append(sort, fmt.Sprintf("%s %s", value, "ASC"))
			}
		}
	}

	return strings.Join(sort, ",")
}

func (p paginator) Valid() (bool, validator.Errors) {
	v := validator.New()
	v.Check(p.page > 0, "page", "must be greater than zero")
	v.Check(p.pageSize > 0, "page_size", "must be greater than zero")
	v.Check(p.pageSize <= paginationMaxPageSize, "page_size", "must be a maximum of 1000")
	return v.Valid(), v.Errors
}

type metadata struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageSize     int `json:"page_size,omitempty"`
	FirstPage    int `json:"first_page,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
}

func (p paginator) CalculateMetadata(totalRecords int) metadata {
	if totalRecords == 0 {
		return metadata{}
	}

	return metadata{
		CurrentPage:  p.page,
		PageSize:     p.pageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(p.pageSize))),
		TotalRecords: totalRecords,
	}
}

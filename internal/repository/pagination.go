package repository

import (
	"net/http"
	"strconv"
)

type PaginatedFeedQuery struct {
	Limit  int    `json:"limit" validate:"gte=1,lte=20"`
	Offset int    `json:"offset" validate:"gte=0"`
	Sort   string `json:"sort" validate:"oneof=asc desc"`
}

func (fq *PaginatedFeedQuery) Parse(r *http.Request) (PaginatedFeedQuery, error) {
	qs := r.URL.Query()

	if limit := qs.Get("limit"); limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return *fq, nil
		}
		fq.Limit = l
	}

	if offset := qs.Get("offset"); offset != "" {
		l, err := strconv.Atoi(offset)
		if err != nil {
			return *fq, nil
		}
		fq.Offset = l
	}

	if sort := qs.Get("sort"); sort != "" {
		fq.Sort = sort
	}

	return *fq, nil
}

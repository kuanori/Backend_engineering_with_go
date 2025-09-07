package repository

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

type PaginatedFeedQuery struct {
	Limit  int      `json:"limit" validate:"gte=1,lte=20"`
	Offset int      `json:"offset" validate:"gte=0"`
	Sort   string   `json:"sort" validate:"oneof=asc desc"`
	Tags   []string `json:"tags" validate:"max=5"`
	Search string   `json:"search" validate:"max=100"`
	Since  string   `json:"since"`
	Until  string   `json:"until"`
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

	tags := qs.Get("tags")
	if tags != "" {
		fq.Tags = strings.Split(tags, ",")
	}

	search := qs.Get("search")
	if search != "" {
		fq.Search = search
	}

	since := qs.Get("since")
	if search != "" {
		fq.Since = parseTime(since)
	}

	until := qs.Get("until")
	if search != "" {
		fq.Until = parseTime(until)
	}

	return *fq, nil
}

func parseTime(s string) string {
	t, err := time.Parse(time.DateTime, s)
	if err != nil {
		return ""
	}

	return t.Format(time.DateTime)
}

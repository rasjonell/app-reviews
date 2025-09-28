package dto

import (
	"fmt"
	"net/http"
	"strconv"
)

func parseParams(w http.ResponseWriter, r *http.Request, key string) (string, bool) {
	id := r.PathValue(key)
	if id == "" {
		http.Error(w, fmt.Sprintf("%s is required", key), http.StatusBadRequest)
		return "", false
	}

	return id, true
}

func parsePagination(r *http.Request) (limit, offset int) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 {
		limit = 10
	}
	offset = (page - 1) * limit

	return
}

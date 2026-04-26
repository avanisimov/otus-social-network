package http

import (
	"net/http"
	"strconv"
)

func getIntQuery(r *http.Request, key string, def int) int {
    v := r.URL.Query().Get(key)
    if v == "" {
        return def
    }

    i, err := strconv.Atoi(v)
    if err != nil {
        return def
    }

    return i
}
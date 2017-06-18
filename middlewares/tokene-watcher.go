package midware

import (
	"net/http"
)

func TokenValidater(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Get("bearer")

		inner.ServeHTTP(w, r)
	})
}

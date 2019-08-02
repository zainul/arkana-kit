package middleware

import (
	"context"
	"net/http"

	"github.com/zainul/ark/xrandom"
)

// ContentType is middleware content type
func ContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("request-id") == "" {
			rd, _ := xrandom.GenerateRandomString(20)
			rqid := "http-user-svc" + rd
			r.Header.Set("request-id", rqid)
		}

		if r.Header.Get("x-roundtrip") == "" {
			r.Header.Set("x-roundtrip", "http-user")
		}

		ctx := context.WithValue(r.Context(), "request-id", r.Header.Get("request-id"))
		xround := r.Header.Get("x-roundtrip")
		ctx = context.WithValue(ctx, "x-roundtrip", "->"+xround)

		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r.WithContext(ctx))
		return
	})
}

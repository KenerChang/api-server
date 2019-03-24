package middleware

import (
	"context"
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

const ContextKeyRequestID = "requestID"

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		uuidObj, err := uuid.NewV4()
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		uuidStr := hex.EncodeToString(uuidObj[:])

		ctx = context.WithValue(ctx, ContextKeyRequestID, uuidStr)

		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

package middleware

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const userKey contextKey = "user"

// AuthMiddleware validates the Bearer token and attaches the claims to the context
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Extract the token from the Authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		claims, err := ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Add the claims to the context
		ctx := context.WithValue(r.Context(), userKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserFromContext retrieves the user claims from the context
func GetUserFromContext(ctx context.Context) *Claims {
	if claims, ok := ctx.Value(userKey).(*Claims); ok {
		return claims
	}
	return nil
}

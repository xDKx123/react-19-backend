package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key") // Use a secure secret key!
var RequestIdHeader = http.CanonicalHeaderKey("X-Request-Id")

func WithMethodChech(handler http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		handler(w, r)
	}
}

// Claims defines the custom JWT claims
type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, email string, expiration time.Duration) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken validates the JWT and returns the claims if valid
func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Extract the claims
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func getRequestId(r *http.Request) string {
	requestId := r.Header.Get(RequestIdHeader)
	if requestId == "" {
		requestId = "0"
	}

	return requestId
}

func requestLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		s := time.Now()
		// most of the required data is already available from requests context.
		ctx := r.Context()
		requestID := getRequestId(r)

		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}

		logger := logrus.WithFields(logrus.Fields{
			"request_id": requestID,
		})

		uri := fmt.Sprintf("%s://%s%s/", scheme, r.Host, r.RequestURI)

		// generate these fields separately, as we will only log them once, to reduce both the visual and memory clutter.
		// all request data can be traced using the request id.
		fields := logrus.Fields{
			"http_scheme": scheme,
			"http_proto":  r.Proto,
			"http_method": r.Method,
			"remote_addr": r.RemoteAddr,
			"request_id":  requestID,
			"user_agent":  r.UserAgent(),
			"uri":         uri,
		}

		// log the only-once fields
		logger.WithFields(fields).Info("new http request")

		ctx = context.WithValue(r.Context(), "request_id", requestID)

		// defer the execution of this function until after the wrapper has run, this allows us to calculate the round trip
		// and log it.
		defer func(s time.Time, logger *logrus.Entry) {
			logger.WithFields(logrus.Fields{
				"request_id": requestID,
				"elapsed":    time.Since(s),
			}).Info("http request processed")
		}(s, logger)

		// next middleware
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

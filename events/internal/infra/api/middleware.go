package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"regexp"
	"slices"

	"github.com/golang-jwt/jwt/v5"
)

type OutputAuth struct {
	ID   string `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type ContextKey string

const ContextUserRoleKey ContextKey = "userRole"

func ValidateToken(tokenString string) (*OutputAuth, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	claims := &OutputAuth{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return &OutputAuth{
		ID:   claims.ID,
		Role: claims.Role,
	}, nil
}

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		p := r.URL.Path
		match, _ := regexp.MatchString("/events/docs/*", p)
		if match {
			next.ServeHTTP(w, r)
			return
		}
		auth := r.Header.Get("Authorization")
		if auth == "" {
			message, statusCode := HandlerErrors(errors.New("the authorization is mandatory"))
			w.WriteHeader(statusCode)
			if _, err := w.Write(message); err != nil {
				log.Panicln(err)
			}
			return
		}
		token := auth[len("Bearer "):]
		data, err := ValidateToken(token)
		if err != nil {
			message, statusCode := HandlerErrors(errors.New("unauthorized user"))
			w.WriteHeader(statusCode)
			if _, err := w.Write(message); err != nil {
				log.Panicln(err)
			}
			return
		}
		if !isValidRole(data.Role) {
			message, statusCode := HandlerErrors(errors.New("unauthorized user"))
			w.WriteHeader(statusCode)
			if _, err := w.Write(message); err != nil {
				log.Panicln(err)
			}
			return
		}
		ctx := context.WithValue(r.Context(), ContextUserRoleKey, data.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		perm, ok := r.Context().Value(ContextUserRoleKey).(string)
		if !ok || perm != "ADMIN" {
			message, statusCode := HandlerErrors(errors.New("you don't have permission to access this resource"))
			w.WriteHeader(statusCode)
			if _, err := w.Write(message); err != nil {
				log.Panicln(err)
			}
			return
		}
		next.ServeHTTP(w, r)
	})
}

func isValidRole(role string) bool {
	rules := []string{"ADMIN", "PUBLIC"}
	if role == "" || !slices.Contains(rules, role) {
		return false
	}
	return true
}

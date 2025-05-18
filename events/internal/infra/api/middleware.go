package api

import (
	"context"
	"net/http"
	"slices"

	"github.com/go-chi/jwtauth/v5"
)

type UserClaim struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}

const (
	AdminRole  = "ADMIN"
	PublicRole = "PUBLIC"
)

var ROLESACCEPTED = []string{AdminRole, PublicRole}

type ContextKey string

const ContextUserRoleKey ContextKey = "userClaims"

func WithJWTAuth(tokenAuth *jwtauth.JWTAuth) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			id, okID := claims["id"].(string)
			role, okRole := claims["role"].(string)
			if !okID || !okRole {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}
			userClaims := &UserClaim{
				ID:   id,
				Role: role,
			}
			if userClaims.Role == "" || !slices.Contains(ROLESACCEPTED, userClaims.Role) {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), ContextUserRoleKey, userClaims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetClaims(r *http.Request) *UserClaim {
	return r.Context().Value(ContextUserRoleKey).(*UserClaim)
}

func OnlyAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := GetClaims(r)
		if claims == nil {
			http.Error(w, "Missing JWT claims", http.StatusUnauthorized)
			return
		}

		if claims.Role != AdminRole {
			http.Error(w, "Admin access required", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

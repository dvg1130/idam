package middleware

import (
	"fmt"
	"net/http"
	"tiered-service-backend/internal/auth"
)

func PayloadLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := auth.GetClaimsFromContext(r.Context())
		if claims == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		role, _ := claims["role"].(string)

		// role limits
		var maxAllowed int
		switch role {
		case "user":
			maxAllowed = 10
		case "tier1":
			maxAllowed = 100
		case "tier2":
			maxAllowed = 500
		case "admin":
			maxAllowed = -1 // unlimited
		default:
			http.Error(w, "Unknown role", http.StatusForbidden)
			return
		}

		username := claims["username"].(string)

		//rbac middleware- quieries db for num of entries based on in snakes tables for user
		//checks against roles maxamount then if db entries != maxallowed error else return

		fmt.Println(w, username, maxAllowed)

		next.ServeHTTP(w, r)
	})
}

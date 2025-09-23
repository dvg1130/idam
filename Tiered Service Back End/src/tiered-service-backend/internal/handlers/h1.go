package handlers

import (
	"fmt"
	"net/http"
	"tiered-service-backend/internal/auth"
	"tiered-service-backend/internal/server"
)

func helper1(_ *server.Server, w http.ResponseWriter, r *http.Request) {
	claims := auth.GetClaimsFromContext(r.Context())
	if claims == nil {
		http.Error(w, "no claims found", http.StatusUnauthorized)
		return

	}
	username := claims["username"].(string)
	// role := claims["role"].(string)

	fmt.Fprintf(w, "%s, your upload was successful", username)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

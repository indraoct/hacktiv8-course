package middleware

import (
	"context"
	"hacktiv8-course/final_assignment/commons/jwt"
	"hacktiv8-course/final_assignment/internal/repositories"
	"net/http"
)

type TokenAuth struct {
	RepoAuth repositories.RepositoryImpl
	JwtPkg   *jwt.JwtPkg
}

func FirstTime(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Create a context with some value
		ctx := context.WithValue(r.Context(), "konco", "octama Indra")

		next.ServeHTTP(w, r.WithContext(ctx)) // call ServeHTTP on the original handler

	})
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Create a context with some value
		ctx := context.WithValue(r.Context(), "kunci", "indra octama")

		next.ServeHTTP(w, r.WithContext(ctx)) // call ServeHTTP on the original handler

	})
}

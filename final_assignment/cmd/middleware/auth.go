package middleware

import (
	"context"
	"hacktiv8-course/final_assignment/commons/constants"
	"hacktiv8-course/final_assignment/commons/jwt"
	"hacktiv8-course/final_assignment/internal/entities"
	"hacktiv8-course/final_assignment/internal/repositories"
	"net/http"
	"strings"
)

type TokenAuth struct {
	RepoAuth repositories.RepositoryImpl
	JwtPkg   *jwt.JwtPkg
}

func NewAuth(repoAuth repositories.RepositoryImpl, jwtPkg *jwt.JwtPkg) TokenAuth {
	return TokenAuth{
		RepoAuth: repoAuth,
		JwtPkg:   jwtPkg,
	}
}

func FirstTime(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Create a context with some value
		ctx := context.WithValue(r.Context(), "konco", "octama Indra")

		next.ServeHTTP(w, r.WithContext(ctx)) // call ServeHTTP on the original handler

	})
}

func (t TokenAuth) Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			token   string
			user    entities.User
			userPtr *entities.User
			err     error
		)

		token = strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)
		user, err = t.JwtPkg.ValidateToken(token)
		if err != nil {
			entities.MessageResponse(err, http.StatusUnauthorized, "", w)
			return
		}

		userPtr, err = t.RepoAuth.GetUserByID(r.Context(), user.ID)
		if err != nil {
			entities.MessageResponse(err, http.StatusUnauthorized, "", w)
			return
		}

		// inject data login to context
		ctx := context.WithValue(r.Context(), constants.KeyUserLogin, *userPtr)

		next.ServeHTTP(w, r.WithContext(ctx)) // call ServeHTTP on the original handler

	})
}

func GetAuth(ctx context.Context) (user entities.User) {
	user = ctx.Value(constants.KeyUserLogin).(entities.User)
	return
}

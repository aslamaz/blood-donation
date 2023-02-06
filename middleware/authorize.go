package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/aslamaz/blood-donation/constant"
	"github.com/aslamaz/blood-donation/repository"
	"github.com/aslamaz/blood-donation/response"
	"github.com/aslamaz/blood-donation/shared"
	"github.com/golang-jwt/jwt/v4"
)

func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")
		token = strings.TrimPrefix(token, "Bearer ")

		t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(constant.JwtSigningKey), nil
		})

		if err != nil {
			shared.SendJson(w, http.StatusForbidden, &response.Response{
				Error: "invalid session",
			})
			return
		}
		if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
			idRaw := claims["id"].(float64) //Type assertion interface{} -> type (float64)
			id := int(idRaw)
			user, err := repository.GetUserById(id)
			if err != nil {
				fmt.Println(err)
				shared.SendJson(w, http.StatusInternalServerError, &response.Response{
					Error: "internal server errorr",
				})
				return
			}
			if user == nil {
				shared.SendJson(w, http.StatusForbidden, &response.Response{
					Error: "invalid session",
				})
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))

		} else {
			shared.SendJson(w, http.StatusForbidden, &response.Response{
				Error: "invalid session",
			})
		}

	})
}

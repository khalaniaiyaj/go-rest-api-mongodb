package middlewares

import (
    "github.com/dgrijalva/jwt-go"
	"strings"
	"fmt"
	"encoding/json"
	"github.com/gorilla/context"
	"net/http"
	. "github.com/user/golang-new/dao"
)

type Middleware struct {}

type Exception struct {
    Message string `json:"message"`
}

var dao = Dao{}

func (m *Middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		 authorizationHeader := req.Header.Get("Authorization")
        if authorizationHeader != "" {
            bearerToken := strings.Split(authorizationHeader, " ")
            if len(bearerToken) == 2 {
                token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
                    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                        return nil, fmt.Errorf("There was an error")
                    }
                    return []byte("secret"), nil
                })
                if error != nil {
                    json.NewEncoder(res).Encode(Exception{Message: "Token is Invalid"})
                    return
                }
                if token.Valid {
                    context.Set(req, "token", token.Claims)
                    next.ServeHTTP(res, req)
                } else {
                    json.NewEncoder(res).Encode(Exception{Message: "Invalid authorization token"})
                }
            }
        } else {
            json.NewEncoder(res).Encode(Exception{Message: "An authorization header is required"})
        }
	})
}

package controller

import (
	"errors"
	"context"
	"fmt"
	"github.com/engajerest/auth/Models/users"
	"github.com/engajerest/auth/logger"
	"github.com/engajerest/auth/utils/Errors"
	"github.com/engajerest/auth/utils/accesstoken"
	"github.com/spf13/viper"
	"net/http"
	"os"
)



func Middleware(key string) func(http.Handler) http.Handler {
	// fmt.Println("1")
	return func(next http.Handler) http.Handler {
		// fmt.Println("2")
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				// fmt.Println("3")
				header := r.Header.Get("token")

				// Allow unauthenticated users in
				if header == "" {
					// http.Error(w, "token null", http.StatusForbidden)
					next.ServeHTTP(w, r)
					return
				}
				// fmt.Println("4")
				//validate jwt token
				tokenStr := header
				userid,configid, err := accesstoken.ParseToken(tokenStr)
				// fmt.Println("5")
				if err != nil {
					http.Error(w, "session expired login again", http.StatusForbidden)
					logger.Error("session token expired", err)
					return
				}
				print("configid=",int(configid))
				if int(configid)==1{
					data1 := users.User{}
					user, status, _ := data1.UserAuthentication(int64(userid))
					print(status)
					print("testing")
					// if usererr !=nil {
					// 	print("st1")
					// 	http.Error(w, "no user found", http.StatusForbidden)
					// 	return
					// }
	
					// put it in context
					ctx := context.WithValue(r.Context(), key, user)
	
					// and call the next with our new context
					r = r.WithContext(ctx)
					next.ServeHTTP(w, r)
				}else{
					print("configid>1")
				}
				

			})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) (*users.User, *Errors.RestError) {
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	userCtxKey := viper.GetString("APP.USER_CONTEXT_KEY")
	print(userCtxKey)
	noUserFoundError := errors.New("no user found")
	if ctx.Value(userCtxKey) == nil {
		return nil, &Errors.RestError{
			Error:   noUserFoundError,
			Message: "no data",
			Code:    http.StatusUnauthorized,
		}
	}
	user, ok := ctx.Value(userCtxKey).(*users.User)
	if !ok || user.ID == 0 {
		return nil, &Errors.RestError{
			Error:   noUserFoundError,
			Message: "no data",
			Code:    http.StatusUnauthorized,
		}
	}
	return user, nil
}

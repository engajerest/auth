package controller

import (
	"github.com/engajerest/auth/Models/users"
	"github.com/engajerest/auth/logger"
	"github.com/engajerest/auth/utils/Errors"
	"github.com/engajerest/auth/utils/accesstoken"
	"errors"

	"context"

	"net/http"

	// "github.com/99designs/gqlgen/graphql"
	// "github.com/vektah/gqlparser/v2/gqlerror"
	// "strconv"
)


var userCtxKey = "usercontextkey"

func Middleware() func(http.Handler) http.Handler {
// fmt.Println("1")
	return func(next http.Handler) http.Handler {
		// fmt.Println("2")
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
			// fmt.Println("3")
			header := r.Header.Get("token")

			// Allow unauthenticated users in
			if header == "" {
			
				next.ServeHTTP(w, r)
				return
			}
			// fmt.Println("4")
			//validate jwt token
			tokenStr := header
		   userid, err := accesstoken.ParseToken(tokenStr)
			// fmt.Println("5")
			if err != nil {
				http.Error(w, "session expired login again", http.StatusForbidden)
				logger.Error("session token expired",err)
				return
			}
			// create user and check if user exists in db
			data1 := users.User{}

	user,status,_ := data1.UserAuthentication(int64(userid))
	print(status)
	print("testing")
			// if usererr !=nil {
			// 	print("st1")
			// 	http.Error(w, "no user found", http.StatusForbidden)
			// 	return
			// }


	// put it in context
	ctx := context.WithValue(r.Context(), userCtxKey, user)

	// and call the next with our new context
	r = r.WithContext(ctx)
	next.ServeHTTP(w, r)
  
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) (*users.User,*Errors.RestError) {
	noUserFoundError:=errors.New("no user found")
	if ctx.Value(userCtxKey) ==nil{
	return nil,	&Errors.RestError{
		Error: noUserFoundError,
		Message: "no data",
		Code: http.StatusBadRequest,
	}
	}
	user, ok := ctx.Value(userCtxKey).(*users.User)
	if !ok || user.ID ==0{
		return nil,&Errors.RestError{
			Error: noUserFoundError,
			Message: "no data",
			Code: http.StatusBadRequest,
		}
	}
return user,nil
}
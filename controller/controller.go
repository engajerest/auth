package controller

import (
	// "bitbucket/engaje_rest_api/engaje-auth/graph"
	"github.com/engajerest/auth/Models/users"
	"github.com/engajerest/auth/utils/accesstoken"
	"fmt"
	"net/http"

	// "bitbucket/engaje_rest_api/engaje-auth/jwt"

	// "github.com/99designs/gqlgen/example/starwars/generated"
	// "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

// var userCtxKey = &contextKey{"user"}

// type contextKey struct {
// 	name string

// }
func PlaygroundHandlers() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}

}
// func GraphHandler() gin.HandlerFunc {
// 	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
// 	return func(c *gin.Context) {
// 		srv.ServeHTTP(c.Writer, c.Request)
// 	}
// }

// func(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"token data": c.Request.Header["Token"],
// 	})
// }

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("tkn1")
		token := c.Request.Header.Get("token")
		fmt.Println("tkn2")
		print(token)
		if token == "" {
			c.JSON(http.StatusUnauthorized, "token null")
			c.Abort()
			return
		}
		fmt.Println("tkn3")
		userId, err := accesstoken.ParseToken(token)
		fmt.Println("5")
		if err != nil {
			c.JSON(http.StatusUnauthorized, "token denied")
			c.Abort()
			return
		}
		fmt.Println("tkn4")
		id := int(userId)
		// create user and check if user exists in db
		data1 := users.User{}

		user,err:= data1.GetByUserId(int64(id))
		if err != nil {
			c.Next()
			return
		}
		print(user.CreatedDate)
				// put it in context
		// ctx := context.WithValue(c.Request.Context(), userCtxKey, user)
	
		// 	// and call the next with our new context
		// 	r= c.Request.WithContext(ctx)
			
		// 		c.Next()
			
				
	}
}
func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	//c.Abort(code)
}
// func ForContext(ctx context.Context) *users.User {
// 	raw, _ := ctx.Value(userCtxKey).(*users.User)
// 	return raw
// }

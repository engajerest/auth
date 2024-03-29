package controller

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/engajerest/auth/Models/users"
	"github.com/engajerest/auth/graph"
	"github.com/engajerest/auth/graph/generated"
	"github.com/engajerest/auth/utils/accesstoken"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PlaygroundHandlers() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}

}
func GraphHandler() gin.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func TokenAuthMiddleware(contextkey string) gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("token")
		print(token)
		if token == "" {
			c.JSON(http.StatusUnauthorized, "token null")
			c.Abort()

			return
		}

		userId, configid, err := accesstoken.ParseToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, "token denied")
			c.Abort()
			return
		}

		id := int(userId)
		id1 := int(configid)
		print("confiid", id1)
		if id1 == 1 {
			print("configid==1")
			data1 := users.User{}
			user, status, errrr := data1.UserAuthentication(int64(id))
			print(status)
			if errrr != nil {
				c.JSON(http.StatusBadRequest, "user not found")
				c.Abort()
				return
			}
			print(user.CreatedDate)
			ctx := context.WithValue(c.Request.Context(), contextkey, user)
			c.Request = c.Request.WithContext(ctx)
			c.Next()
		} else if id1 == 4 {
			print("configid=4 customer")
			data1 := users.User{}
			user, status, errrr := data1.Customerauthenticate(int64(id))
			print(status)
			if errrr != nil {
				c.JSON(http.StatusBadRequest, "user not found")
				c.Abort()
				return
			}
			print(user.CreatedDate)
			ctx := context.WithValue(c.Request.Context(), contextkey, user)
			c.Request = c.Request.WithContext(ctx)
			c.Next()

		} else if id1 == 7 {
			print("id=7 session based customer")
			var user users.User
			user.ID = id
			user.Configid = id1
			ctx := context.WithValue(c.Request.Context(), contextkey, user)
			c.Request = c.Request.WithContext(ctx)
			c.Next()
		} else {
			print("configid not 4")
			c.JSON(http.StatusUnauthorized, "configid not 4")
			c.Abort()
			return
		}

	}
}

func TokenNoAuthMiddleware(contextkey string) gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("token")

		print(token)
		if token == "" {
			// c.JSON(http.StatusUnauthorized, "token null")
			// c.Abort()
			c.Next()
			return
		}

		userId, configid, err := accesstoken.ParseToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, "token denied")
			c.Abort()
			return
		}

		id := int(userId)
		id1 := int(configid)
		print("confiid", id1)
		if id1 == 1 {
			print("configid==1")
			data1 := users.User{}
			user, status, errrr := data1.UserAuthentication(int64(id))
			print(status)
			if errrr != nil {
				c.JSON(http.StatusBadRequest, "user not found")
				c.Abort()
				return
			}
			print(user.CreatedDate)
			ctx := context.WithValue(c.Request.Context(), contextkey, user)
			c.Request = c.Request.WithContext(ctx)
			c.Next()
		} else if id1 == 4 {
			print("configid=4 customer")
			data1 := users.User{}
			user, status, errrr := data1.Customerauthenticate(int64(id))
			print(status)
			if errrr != nil {
				c.JSON(http.StatusBadRequest, "user not found")
				c.Abort()
				return
			}
			print(user.CreatedDate)
			ctx := context.WithValue(c.Request.Context(), contextkey, user)
			c.Request = c.Request.WithContext(ctx)
			c.Next()

		} else if id1 == 7 {
			print("id=7 session based customer")
			data1 := users.User{}
			user := data1.Getzeroauth(int64(id), int64(id1))
			ctx := context.WithValue(c.Request.Context(), contextkey, user)
			c.Request = c.Request.WithContext(ctx)
			c.Next()
		} else {
			print("configid not 4")
			c.JSON(http.StatusUnauthorized, "configid not 4")
			c.Abort()
			return
		}

	}
}

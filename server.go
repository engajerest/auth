package main

import (
	// "bitbucket/engaje_rest_api/engaje-auth/controller"
	"github.com/engajerest/auth/controller"
	"github.com/engajerest/auth/graph"
	"github.com/engajerest/auth/graph/generated"
	"github.com/engajerest/auth/logger"
	"github.com/engajerest/auth/utils/dbconfig"

	"fmt"
	"log"
	"net/http"

	// "net/smtp"
	"os"

	// "github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"

	"github.com/spf13/viper"
)

func main() {
	
	// Config
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	// Declare var
	defaultPort := viper.GetString("APP.PORT")
	dbName := viper.GetString("APP.DATABASE_NAME")
	password := viper.GetString("APP.DATABASE_PASSWORD")
	userName := viper.GetString("APP.DATABASE_USERNAME")
	_ = viper.GetString("APP.DATABASE_PORT")
	host := viper.GetString("APP.DATABASE_SERVER_HOST")
	userCtxKey:=viper.GetString("APP.USER_CONTEXT_KEY")
	fmt.Println("PORT :", defaultPort)

	router := chi.NewRouter()
	router.Use(controller.Middleware(userCtxKey))
	dbconfig.InitDB(dbName, userName, password, host)
    logger.Info("application started")

	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("Engaje", "/query"))
	router.Handle("/auth", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, router))

}

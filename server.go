package main

import (
	"fmt"
	"log"
	"os"
	"github.com/engajerest/auth/controller"
	"github.com/engajerest/auth/logger"
	"github.com/engajerest/auth/utils/dbconfig"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	print("1")
	// Config
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	print("2")
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	print("3")
	// Declare var
	defaultPort := viper.GetString("APP.PORT")
	dbName := viper.GetString("APP.DATABASE_NAME")
	password := viper.GetString("APP.DATABASE_PASSWORD")
	userName := viper.GetString("APP.DATABASE_USERNAME")
	_ = viper.GetString("APP.DATABASE_PORT")
	host := viper.GetString("APP.DATABASE_SERVER_HOST")
	userCtxKey := viper.GetString("APP.USER_CONTEXT_KEY")
	fmt.Println("PORT :", defaultPort)

	router := gin.Default()

	dbconfig.InitDB(dbName, userName, password, host)
	logger.Info("auth application started")
	router.Use(controller.TokenNoAuthMiddleware(userCtxKey))
	router.GET("/", controller.PlaygroundHandlers())
	router.POST("/auth", controller.GraphHandler())
	router.Run(defaultPort)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)

	// router := chi.NewRouter()
	// router.Use(controller.Middleware(userCtxKey))
	// dbconfig.InitDB(dbName, userName, password, host)
	// logger.Info("application started")

	// server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	// router.Handle("/", playground.Handler("Engaje", "/query"))
	// router.Handle("/auth", server)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	// log.Fatal(http.ListenAndServe(":"+defaultPort, router))

}

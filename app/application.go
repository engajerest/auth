package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/engajerest/auth/controller"
	"github.com/engajerest/auth/utils/dbconfig"

	// "github.com/engajerest/sparkle/controllers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	router = gin.Default()
)

func StartApplication() {
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	defaultPort := viper.GetString("APP.PORT")
	key := viper.GetString("APP.USER_CONTEXT_KEY")
	print("key====",key)
	
	router.Use(Datasource())
	router.Use(controller.TokenNoAuthMiddleware(key))
	fmt.Print("Heloo world")
	
	Mapurls()
	router.Run(defaultPort)

}
func Datasource() gin.HandlerFunc {
	return func(c *gin.Context) {

		metrics := map[string]string{
			"method":   c.Request.Method,
			"endpoint": c.FullPath(),
		}
		met1 := metrics["endpoint"]

		print(met1)
		result := strings.Split(met1, "/")
		print("rawpath===", met1)
		fmt.Println("splitpath===", result)
		var flavour string
		for _ = range result {
			flavour = result[1]
			print("flav==", flavour)
		}

		if flavour == "dev" {

			print("Now its Dev")
			//Config
			viper.SetConfigName("config") // config file name without extension
			viper.SetConfigType("yaml")
			viper.AddConfigPath(".")
			err := viper.ReadInConfig()
			if err != nil {
				fmt.Println("fatal error config file: default \n", err)
				os.Exit(1)
			}

			// Declare var
			defaultPort := viper.GetString("DEV.PORT")
			dbName := viper.GetString("DEV.DATABASE_NAME")
			password := viper.GetString("DEV.DATABASE_PASSWORD")
			userName := viper.GetString("DEV.DATABASE_USERNAME")
			_ = viper.GetString("DEV.DATABASE_PORT")
			host := viper.GetString("DEV.DATABASE_SERVER_HOST")
			UserContextKey := viper.GetString("DEV.USER_CONTEXT_KEY")
			fmt.Println("PORT :", defaultPort)
			fmt.Println("userkey :", UserContextKey)

			dbconfig.InitDB(dbName, userName, password, host)
			print("db configured for Dev")
		} else if flavour == "v1" {
			print("Now its Live")
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
			UserContextKey := viper.GetString("APP.USER_CONTEXT_KEY")
			fmt.Println("PORT :", defaultPort)
			fmt.Println("userkey :", UserContextKey)

			dbconfig.InitDB(dbName, userName, password, host)
			print("db configured for live")
		}
		c.Next()

	}
}

package accesstoken

import (
	"github.com/engajerest/auth/logger"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"github.com/spf13/viper"
	"github.com/dgrijalva/jwt-go"
	"os"
)

// GenerateToken generates a jwt token and assign a username to it's claims and return it
func GenerateToken(userid,configid int) (string, error) {
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	Key := viper.GetString("APP.JWT_SECRET_KEY")
	SecretKey := []byte(Key)

	token := jwt.New(jwt.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)
	/* Set token claims */
	claims["userid"] = userid
	claims["configid"]=configid
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating key")
		return "", err
	}
	
	return tokenString, nil
}

// ParseToken parses a jwt token and returns the username in it's claims
func ParseToken(tokenStr string) (userid,configid float64,Error error) {
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	Key := viper.GetString("APP.JWT_SECRET_KEY")
	SecretKey := []byte(Key)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		
		userid := claims["userid"].(float64)
		configid:=claims["configid"].(float64)
		
		
		var tm time.Time
		switch iat := claims["exp"].(type) {
		case float64:
			tm = time.Unix(int64(iat), 0)
		case json.Number:
			v, _ := iat.Int64()
			tm = time.Unix(v, 0)
		}
	
		fmt.Println(tm)
		logger.Time("expiry time",tm)
		return userid,configid, nil
	} else {
		return 0,0, err
	}
}
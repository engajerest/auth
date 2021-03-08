package contextdata

import (
	"context"

	"github.com/engajerest/auth/Models/users"
	"github.com/engajerest/auth/utils/Errors"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"errors"
	"fmt"
)


func ForAuthContext(ctx context.Context) (*users.User, *Errors.RestError) {
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	userCtxKey := viper.GetString("APP.USER_CONTEXT_KEY")
	noUserFoundError := errors.New("no user found")
	if ctx.Value(userCtxKey) == nil {
		return nil, &Errors.RestError{
			Error:   noUserFoundError,
			Message: "no data",
			Code:    http.StatusBadRequest,
		}
	}
	user, ok := ctx.Value(userCtxKey).(*users.User)
	if !ok || user.ID == 0 {
		return nil, &Errors.RestError{
			Error:   noUserFoundError,
			Message: "no data",
			Code:    http.StatusBadRequest,
		}
	}
	return user, nil
}
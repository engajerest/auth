package Errors

// import (
	
// "net/http"
// )

type WrongUsernameOrPasswordError struct{}
type RestError struct {
    Error   error
    Message string
    Code    int
}
func (m *WrongUsernameOrPasswordError) Error() string {
	return "wrong username or password"
}
// type RestError struct {
// 	Status  bool   `json:"status"`
// 	Message string `json:"message"`
// 	Code    int    `json:"code"`
// 	Error   string `json:"Error"`
// }

// func NewBadRequestError(error string) *RestError{
//       return &RestError{
// 		  Status: false,
// 		  Message: "Internal Error",
// 		  Code: http.StatusBadRequest,
// 		  Error: error,
// 	  }
// }
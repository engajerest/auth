package users

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	CreatedDate string `json:"created"`
	Status      string `json:"status"`
}





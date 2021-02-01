// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type GetUser struct {
	UserID    int    `json:"userId"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	Created   string `json:"created"`
	Status    string `json:"status"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginData struct {
	Status   bool      `json:"status"`
	Code     int       `json:"code"`
	Message  string    `json:"message"`
	UserInfo *UserData `json:"userInfo"`
}

type NewUser struct {
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Mobile      string `json:"Mobile"`
	Roleid      *int   `json:"roleid"`
	Referenceid *int   `json:"referenceid"`
	Locationid  *int   `json:"locationid"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type Reset struct {
	Password string `json:"password"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserCreatedData struct {
	Status   bool      `json:"status"`
	Code     int       `json:"code"`
	Message  string    `json:"message"`
	UserInfo *UserData `json:"userInfo"`
}

type UserData struct {
	UserID      int     `json:"UserId"`
	Tenantid    *int    `json:"Tenantid"`
	Locationid  *int    `json:"Locationid"`
	Moduleid    *int    `json:"Moduleid"`
	Packageid   *int    `json:"Packageid"`
	Firstname   string  `json:"Firstname"`
	Lastname    string  `json:"Lastname"`
	Email       string  `json:"Email"`
	Mobile      string  `json:"Mobile"`
	Token       string  `json:"Token"`
	Modulename  *string `json:"Modulename"`
	Tenantname  *string `json:"Tenantname"`
	CreatedDate string  `json:"CreatedDate"`
	Status      string  `json:"Status"`
}

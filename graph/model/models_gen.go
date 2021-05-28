// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type GetUser struct {
	Userid         int    `json:"Userid"`
	Firstname      string `json:"Firstname"`
	Lastname       string `json:"Lastname"`
	Mobile         string `json:"Mobile"`
	Email          string `json:"Email"`
	Profileimage   string `json:"Profileimage"`
	Currencysymbol string `json:"Currencysymbol"`
	Currencycode   string `json:"Currencycode"`
	Countrycode    string `json:"Countrycode"`
	Created        string `json:"Created"`
	Status         string `json:"Status"`
}

type Login struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Tenanttoken string `json:"tenanttoken"`
	Devicetype  string `json:"devicetype"`
}

type LoginData struct {
	Status       bool          `json:"status"`
	Code         int           `json:"code"`
	Message      string        `json:"message"`
	UserInfo     *UserData1    `json:"userInfo"`
	Tenantinfo   []*Tenantdata `json:"tenantinfo"`
	Locationinfo []*Location   `json:"locationinfo"`
}

type NewUser struct {
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Mobile         string `json:"Mobile"`
	Roleid         int    `json:"roleid"`
	Referenceid    int    `json:"referenceid"`
	Locationid     int    `json:"locationid"`
	Configid       int    `json:"configid"`
	Countrycode    string `json:"countrycode"`
	Currencycode   string `json:"currencycode"`
	Currencysymbol string `json:"currencysymbol"`
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
	UserID         int     `json:"UserId"`
	Firstname      string  `json:"Firstname"`
	Lastname       string  `json:"Lastname"`
	Email          string  `json:"Email"`
	Mobile         string  `json:"Mobile"`
	Profileimage   string  `json:"Profileimage"`
	Currencysymbol string  `json:"Currencysymbol"`
	Currencycode   string  `json:"Currencycode"`
	Countrycode    string  `json:"Countrycode"`
	Roleid         *int    `json:"Roleid"`
	Configid       *int    `json:"Configid"`
	Token          string  `json:"Token"`
	Tenantid       *int    `json:"Tenantid"`
	Tenantname     *string `json:"Tenantname"`
	Tenantimageurl *string `json:"Tenantimageurl"`
	Locationid     *int    `json:"Locationid"`
	Opentime       *string `json:"Opentime"`
	Closetime      *string `json:"Closetime"`
	CreatedDate    string  `json:"CreatedDate"`
	Status         string  `json:"Status"`
}

type UserData1 struct {
	UserID             int     `json:"UserId"`
	Firstname          string  `json:"Firstname"`
	Lastname           string  `json:"Lastname"`
	Email              string  `json:"Email"`
	Mobile             string  `json:"Mobile"`
	Profileimage       string  `json:"Profileimage"`
	Usercurrencysymbol string  `json:"Usercurrencysymbol"`
	Usercurrencycode   string  `json:"Usercurrencycode"`
	Usercountrycode    string  `json:"Usercountrycode"`
	Roleid             *int    `json:"Roleid"`
	Configid           *int    `json:"Configid"`
	Token              string  `json:"Token"`
	Tenantid           *int    `json:"Tenantid"`
	Tenantname         *string `json:"Tenantname"`
	Tenantimageurl     *string `json:"Tenantimageurl"`
	Tenantaccid        string  `json:"Tenantaccid"`
	Countrycode        string  `json:"Countrycode"`
	Currencyid         int     `json:"Currencyid"`
	Currencycode       string  `json:"Currencycode"`
	Currencysymbol     string  `json:"Currencysymbol"`
	CreatedDate        string  `json:"CreatedDate"`
	Status             string  `json:"Status"`
	Devicetype         string  `json:"Devicetype"`
}

type Location struct {
	Locationid   int    `json:"Locationid"`
	Tenantid     int    `json:"Tenantid"`
	Locationname string `json:"Locationname"`
	Email        string `json:"Email"`
	Contactno    string `json:"Contactno"`
	Address      string `json:"Address"`
	Suburb       string `json:"Suburb"`
	City         string `json:"City"`
	State        string `json:"State"`
	Postcode     string `json:"Postcode"`
	Latitude     string `json:"Latitude"`
	Longitude    string `json:"Longitude"`
	Opentime     string `json:"Opentime"`
	Closetime    string `json:"Closetime"`
}

type Tenantdata struct {
	Subscriptionid       int     `json:"Subscriptionid"`
	Packageid            int     `json:"Packageid"`
	Packagename          string  `json:"Packagename"`
	Moduleid             int     `json:"Moduleid"`
	Featureid            int     `json:"Featureid"`
	Modulename           string  `json:"Modulename"`
	Validitydate         string  `json:"Validitydate"`
	Validity             bool    `json:"Validity"`
	Categoryid           int     `json:"Categoryid"`
	Subcategoryid        int     `json:"Subcategoryid"`
	Iconurl              string  `json:"Iconurl"`
	Logourl              string  `json:"Logourl"`
	Paymentstatus        bool    `json:"Paymentstatus"`
	Subscriptionaccid    string  `json:"Subscriptionaccid"`
	Subscriptionmethodid string  `json:"Subscriptionmethodid"`
	Taxamount            float64 `json:"Taxamount"`
	Totalamount          float64 `json:"Totalamount"`
}

type Updateddata struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Userupdateinput struct {
	Userid       int    `json:"userid"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Email        string `json:"email"`
	Contactno    string `json:"contactno"`
	Profileimage string `json:"profileimage"`
}

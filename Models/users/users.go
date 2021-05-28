package users

type User struct {
	ID                 int    `json:"id"`
	FirstName          string `json:"firstname"`
	LastName           string `json:"lastname"`
	Password           string `json:"password"`
	Email              string `json:"email"`
	Mobile             string `json:"mobile"`
	Profileimage       string `json:"profileimage"`
	CreatedDate        string `json:"created"`
	Status             string `json:"status"`
	Roleid             int    `json:"roleid"`
	Configid           int    `json:"configid"`
	Referenceid        int    `json:"referenceid"`
	LocationId         int    `json:"locationid"`
	Moduleid           int    `json:"moduleid"`
	Packageid          int    `json:"packageid"`
	Modulename         string `json:"modulename"`
	Tenantname         string `json:"tenantname"`
	Tenantimage        string `json:"tenantimage"`
	Opentime           string `json:"opentime"`
	Closetime          string `json:"closetime"`
	From               string `json:"from"`
	Tenantaccid        string `json:"tenantaccid"`
	Countrycode        string `json:"countrycode"`
	Currencyid         int    `json:"currencyid"`
	Currencysymbol     string `json:"currencysymbol"`
	CurrencyCode       string `json:"currencycode"`
	Devicetype         string `json:"devicetype"`
	Usercountrycode    string `json:"usercountrycode"`
	Usercurrencysymbol string `json:"usercurrencysymbol"`
	UsercurrencyCode   string `json:"usercurrencycode"`
}
type Tenant struct {
	Subscriptionid       int     `json:"subscriptonid"`
	Moduleid             int     `json:"moduleid"`
	Featureid int `json:"featureid"`
	Packageid            int     `json:"packageid"`
	Modulename           string  `json:"modulename"`
	Iconurl              string  `json:"iconurl"`
	Logourl              string  `json:"imageurl"`
	Packagename          string  `json:"packagename"`
	Validiydate          string  `json:"validitydate"`
	Validity             bool    `json:"validity"`
	Subcategoryid        int     `json:"subcategoryid"`
	Categoryid           int     `json:"categoryid"`
	Paymentstatus        bool    `json:"paymentstatus"`
	Subscriptionmethodid string  `json:"subscriptionmethodid"`
	Subscriptionaccid    string  `json:"subscriptionaccid"`
	Taxamount            float64 `json:"taxamount"`
	Totalamount          float64 `json:"totalamount"`
}
type Location struct {
	LocationId   int    `json:"locationid"`
	Tenantid     int    `json:"tenantid"`
	Locationname string `json:"locationname"`
	Email        string `json:"email"`
	Contactno    string `json:"contactno"`
	Address      string `json:"address"`
	Suburb       string `json:"suburb"`
	City         string `json:"city"`
	State        string `json:"state"`
	Postcode     string `json:"postcode"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	Opentime     string `json:"opentime"`
	Closetime    string `json:"closetime"`
}

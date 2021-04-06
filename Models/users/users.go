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
	Roleid      int    `json:"roleid"`
	Configid    int    `json:"configid"`
	Referenceid int    `json:"referenceid"`
	LocationId  int    `json:"locationid"`
	Moduleid    int    `json:"moduleid"`
	Packageid   int    `json:"packageid"`
	Modulename  string `json:"modulename"`
	Tenantname  string `json:"tenantname"`
	Tenantimage string `json:"tenantimage"`
	Opentime    string `json:"opentime"`
	Closetime   string `json:"closetime"`
	From string `json:"from"`
}
type Tenant struct {
	Subscriptionid int `json:"subscriptonid"`
	Moduleid    int    `json:"moduleid"`
	Packageid   int    `json:"packageid"`
	Modulename  string `json:"modulename"`
	Iconurl string `json:"iconurl"`
	Logourl string `json:"imageurl"`
	Packagename  string `json:"packagename"`
   Validiydate string `json:"validitydate"`
   Validity bool `json:"validity"`
   Subcategoryid int `json:"subcategoryid"`
	Categoryid int `json:"categoryid"`
}

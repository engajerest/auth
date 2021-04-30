package users

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/engajerest/auth/logger"
	"github.com/engajerest/auth/utils/Errors"
	"github.com/engajerest/auth/utils/accesstoken"
	database "github.com/engajerest/auth/utils/dbconfig"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const (
	createUserQuery            = "INSERT INTO app_users (authname,password,hashsalt,contactno,roleid,configid) VALUES(?,?,?,?,?,?)"
	createUsernopassword       = "INSERT INTO app_users (authname,contactno,roleid,configid) VALUES(?,?,?,?)"
	insertUsertoProfileQuery   = "INSERT INTO app_userprofiles (userid,firstname,lastname,email,contactno) VALUES(?,?,?,?,?)"
	getUseridByNameQuery       = "select user_id, email, mobile,status,created_date from engaje_users WHERE user_name = ?"
	authenticationQuery        = "SELECT userid,password,hashsalt,IFNULL(configid,0) AS configid FROM app_users WHERE authname=? OR contactno=? AND password=? "
	usersGetAllQuery           = "select userid, firstname,lastname, contactno,email,IFNULL(profileimage,'') AS profileimage, status,created from app_userprofiles"
	getUserByidQuery           = "select userid, firstname,lastname,contactno,email,status,created from app_userprofiles WHERE userid=?"
	resetPasswordQuery         = "UPDATE app_users SET password=? ,hashsalt=?  WHERE userid = ?"
	insertTokentoSessionQuery  = "INSERT INTO app_session (userid,sessionname,sessiondate,sessionexpiry) VALUES(?,?,?,?)"
	checkUseridinSessionQuery  = "select userid from app_session WHERE userid= ?"
	userAuthentication         = "SELECT a.userid,a.roleid,a.configid,b.firstname,b.lastname,b.email,b.contactno,IFNULL(b.profileimage,'') AS profileimage,b.status,b.created FROM app_users a, app_userprofiles b WHERE a.userid=b.userid AND a.status ='Active' AND a.userid=?"
	loginResponseQueryByUserid = "SELECT a.userid,b.firstname,b.lastname,b.contactno,b.email,IFNULL(b.userlocationid,0) AS userlocationid,b.status,b.created, IFNULL(c.tenantid,0) AS tenantid,IFNULL(c.tenantname,'') AS tenantname, IFNULL(d.packageid,0) AS packageid, IFNULL(d.moduleid,0) AS moduleid, IFNULL(e.modulename,'') AS modulename, IFNULL(f.opentime,'') AS opentime,IFNULL(f.closetime,'') AS closetime  FROM app_users a INNER JOIN app_userprofiles b ON a.userid = b.userid LEFT OUTER JOIN tenants c ON a.referenceid=c.tenantid LEFT OUTER JOIN tenantsubscription d ON c.tenantid=d.tenantid LEFT OUTER JOIN app_module e ON d.moduleid=e.moduleid  LEFT OUTER JOIN tenantlocations f ON c.tenantid=f.tenantid WHERE a.userid=?"
	updateenanttoken           = "UPDATE tenants SET tenanttoken=? WHERE tenantid=?"
	checkauthname              = "SELECT userid,IFNULL(configid,0) AS configid FROM app_users WHERE authname= ? OR contactno=?"
	getCustomerByid            = "SELECT customerid,firstname,lastname,contactno,email,IFNULL(configid,0) AS configid  FROM customers WHERE customerid=?"
	updateappuser              = "UPDATE app_users SET authname=? , contactno=? WHERE userid=?"
	updateuserprofile          = "UPDATE app_userprofiles SET firstname=?,lastname=?,email=?,contactno=?,profileimage=? WHERE userid=?"
	loginuserresponse          = "SELECT a.userid,a.authname,a.contactno,a.roleid,a.configid,a.status,a.created,b.firstname,b.lastname,IFNULL(b.profileimage,'') AS profileimage,IFNULL(c.tenantid,0) AS tenantid,IFNULL(c.tenantname,'') AS tenantname,IFNULL(c.tenantimage,'') AS tenantimage FROM app_users a INNER JOIN app_userprofiles b ON a.userid = b.userid LEFT OUTER JOIN tenants c ON a.referenceid=c.tenantid  WHERE   a.userid=?"
	logintenantresponse        = "SELECT a.subscriptionid,a.packageid,a.moduleid,a.categoryid,a.subcategoryid,IFNULL(a.validitydate,'') AS validitydate,IF(a.validitydate>=DATE(NOW()), true, false) AS validity,a.paymentstatus,IFNULL(a.subscriptionaccid,'') AS subscriptionaccid,IFNULL(a.subscriptionmethodid,'') AS subscriptionmethodid,b.modulename,IFNULL(b.logourl,'') AS logourl,IFNULL(b.iconurl,'') AS iconurl FROM tenantsubscription a,app_module b WHERE a.moduleid=b.moduleid  AND  tenantid=? ORDER BY a.subscriptionid ASC "
	loginlocationresponse      = "SELECT locationid,tenantid,locationname,email,contactno,address,city,state,postcode,IFNULL(latitude,'') AS latitude,IFNULL(longitude,'') AS longitude,IFNULL(opentime,'') AS opentime,IFNULL(closetime,'') AS closetime  FROM tenantlocations  WHERE tenantid=?"
)

func (user *User) Create() (int64, error) {
	fmt.Println("0")
	statement, err := database.Db.Prepare(createUserQuery)
	fmt.Println("1")
	if err != nil {
		// log.Fatal(err)
		return 0, err
	}
	defer statement.Close()
	hashedPassword, err := HashPassword(user.Password)
	fmt.Println("2")

	res, err1 := statement.Exec(&user.Email, &user.Password, &hashedPassword, &user.Mobile, &user.Roleid, &user.Configid)
	if err1 != nil {

		// log.Fatal(err1)
		return 0, err1

	}
	id, err2 := res.LastInsertId()
	if err2 != nil {
		// log.Fatal("Error:", err2.Error())
		return 0, err2
	}
	log.Print("Row inserted!")
	return id, nil
}
func (user *User) Createwithoutpassword() (int64, error) {
	fmt.Println("nopass")
	statement, err := database.Db.Prepare(createUsernopassword)
	fmt.Println("1")
	if err != nil {
		print(err)
		return 0, err

	}
	defer statement.Close()

	fmt.Println("2")

	res, err1 := statement.Exec(&user.Email, &user.Mobile, &user.Roleid, &user.Configid)
	if err1 != nil {

		fmt.Println(err1)

		return 0, err1

	}
	id, err2 := res.LastInsertId()
	if err2 != nil {
		log.Fatal("Error:", err2.Error())
		return 0, err
	}
	log.Print("Row inserted!")
	return id, nil
}
func (user *User) InsertUserintoProfile() int64 {
	statement, err := database.Db.Prepare(insertUsertoProfileQuery)
	print(statement)

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	res, err := statement.Exec(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Mobile)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {

	fmt.Println("4")
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println("5")
	if err != nil {
		fmt.Println("false")
		log.Println(err)
		return false
	}
	fmt.Println("true")
	return true
}
func (user *User) GetUserIdByUsername(username string) (int, error) {
	statement, err := database.Db.Prepare(getUseridByNameQuery)
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(username)

	var Id int
	var Email string
	var Mobile string
	var Status string
	var CreatedDate string
	err = row.Scan(&Id, &Email, &Mobile, &Status, &CreatedDate)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}
	fmt.Println(user)
	user.ID = Id
	user.FirstName = username
	user.Email = Email
	user.CreatedDate = CreatedDate
	user.Mobile = Mobile
	user.Status = Status

	return Id, nil
}
func (user *User) Authenticate() bool {
	var data User
	statement, err := database.Db.Prepare(authenticationQuery)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("1")
	row := statement.QueryRow(user.FirstName, user.FirstName, user.Password)

	var hashedPassword string

	err = row.Scan(&data.ID, &data.Password, &hashedPassword, &data.Configid)
	print(err)
	fmt.Println("2")
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}
	// fmt.Println(user)
	user.ID = data.ID
	user.Configid = data.Configid
	return CheckPasswordHash(user.Password, hashedPassword)
}

func (user *User) Checkauthname() bool {
	var data User
	statement, err := database.Db.Prepare(checkauthname)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("1")
	row := statement.QueryRow(user.FirstName, user.FirstName)
	err = row.Scan(&data.ID, &data.Configid)
	print(err)
	fmt.Println("2")
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			return false

		}
	}
	// fmt.Println(user)
	user.ID = data.ID
	user.Configid = data.Configid
	return true
}

func GetAllUsers() []User {
	stmt, err := database.Db.Prepare(usersGetAllQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Mobile, &user.Email, &user.Profileimage, &user.Status, &user.CreatedDate)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return users
}
func Tenantresponse(userid int) []Tenant {
	stmt, err := database.Db.Prepare(logintenantresponse)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(userid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var list []Tenant
	for rows.Next() {
		var t Tenant
		err := rows.Scan(&t.Subscriptionid, &t.Packageid, &t.Moduleid, &t.Categoryid, &t.Subcategoryid, &t.Validiydate, &t.Validity, &t.Paymentstatus,&t.Subscriptionaccid,&t.Subscriptionmethodid, &t.Modulename, &t.Logourl, &t.Iconurl)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, t)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return list
}
func Locationresponse(userid int) []Location {
	stmt, err := database.Db.Prepare(loginlocationresponse)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(userid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var list []Location
	for rows.Next() {
		var t Location
		err := rows.Scan(&t.LocationId,&t.Tenantid,&t.Locationname,&t.Email,&t.Contactno,&t.Address,&t.City,
		&t.State,&t.Postcode,&t.Latitude,&t.Longitude,&t.Opentime,&t.Closetime)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, t)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return list
}
func (user *User) GetByUserId(id int64) (*User, error) {

	fmt.Println("enrty in getbyid")
	print(id)
	var data User
	stmt, err := database.Db.Prepare(getUserByidQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	// print(row)
	err = row.Scan(&data.ID, &data.FirstName, &data.LastName, &data.Mobile, &data.Email, &data.Status, &data.CreatedDate)
	print(err)
	fmt.Println("2")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows found")
		} else {
			log.Fatal(err)
		}
	}
	// fmt.Println(user)
	user.ID = data.ID
	user.FirstName = data.FirstName
	user.LastName = data.LastName
	user.Email = data.Email
	user.CreatedDate = data.CreatedDate
	user.Mobile = data.Mobile
	user.Status = data.Status
	print(user.ID)
	// print(user.FirstName)
	// print(user.LastName)
	// print(user.Email)

	fmt.Println("completed")
	return &data, nil
}
func (user *User) LoginResponse(id int64) (*User, error) {

	fmt.Println("enrty in login response")
	print(id)
	var data User
	stmt, err := database.Db.Prepare(loginuserresponse)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	// print(row)
	err = row.Scan(&data.ID, &data.Email, &data.Mobile, &data.Roleid, &data.Configid, &data.Status,
		&data.CreatedDate, &data.FirstName, &data.LastName, &data.Profileimage, &data.Referenceid, &data.Tenantname, &data.Tenantimage)
	print(err)
	fmt.Println("2")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows found")
		} else {
			log.Fatal(err)
		}
	}
	// fmt.Println(user)
	user.ID = data.ID
	user.FirstName = data.FirstName
	user.LastName = data.LastName
	user.Email = data.Email
	user.CreatedDate = data.CreatedDate
	user.Mobile = data.Mobile
	user.Status = data.Status
	user.Referenceid = data.Referenceid
	user.Roleid = data.Roleid
	user.Configid = data.Configid
	// user.Modulename = data.Modulename
	user.Tenantname = data.Tenantname
	// user.LocationId = data.LocationId
	// user.Closetime = data.Closetime
	// user.Opentime = data.Opentime
	user.Tenantimage = data.Tenantimage
	user.Profileimage = data.Profileimage
	print(user.ID)
	// print(user.FirstName)
	// print(user.LastName)
	// print(user.Email)

	fmt.Println("completed")
	return &data, nil
}

func (user *User) UserAuthentication(id int64) (*User, bool, error) {
	fmt.Println("enrty in auth")
	print(id)
	var data User
	stmt, err := database.Db.Prepare(userAuthentication)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	err = row.Scan(&data.ID,&data.Roleid,&data.Configid, &data.FirstName, &data.LastName, &data.Email, &data.Mobile, &data.Profileimage, &data.Status, &data.CreatedDate)
	print(err)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows found")
			data1 := Errors.RestError{}
			data1.Error = err
			return &data, false, err
		} else {
			log.Fatal(err)
			fmt.Println("nodata")
			var data1 *Errors.RestError
			data1.Error = err
			return &data, false, err
		}

	}
	data.From = "USER"
	// user.Check=true
	return &data, true, err
}
func (c *User) Customerauthenticate(id int64) (*User, bool, error) {

	fmt.Println("enrty in customergetbyid")
	print(id)
	var data User
	stmt, err := database.Db.Prepare(getCustomerByid)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	// print(row)
	err = row.Scan(&data.ID, &data.FirstName, &data.LastName, &data.Mobile, &data.Email, &data.Configid)
	print(err)
	fmt.Println("2")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows found")
			data1 := Errors.RestError{}
			data1.Error = err
			return &data, false, err
		} else {
			log.Fatal(err)
			fmt.Println("nodata")
			var data1 *Errors.RestError
			data1.Error = err
			return &data, false, err
		}

	}
	data.From = "CUSTOMER"
	fmt.Println("completed")
	return &data, true, nil
}

func (user *User) ResetPassword() bool {
	stmt, err := database.Db.Prepare(resetPasswordQuery)
	if err != nil {
		log.Fatal(err)
	}
	hashedPassword, err := HashPassword(user.Password)
	_, err = stmt.Exec(user.Password, hashedPassword, user.ID)
	if err != nil {
		log.Fatal(err)
	}
	return true

}
func (user *User) InsertToken(token string) {
	// fmt.Println("0")
	statement, err := database.Db.Prepare(insertTokentoSessionQuery)
	// fmt.Println("1")
	defer statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	sessiondate := time.Now()
	sessionexpiry := time.Hour.Hours()
	res, err := statement.Exec(&user.ID, token, sessiondate, sessionexpiry)
	if err != nil {
		// print(token)
		logger.Error("database error", err)
		log.Fatal(err)
	}
	_, err = res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("token inserted!")

}

func CheckStatusinSession(id int64) bool {
	stmt, err := database.Db.Prepare(checkUseridinSessionQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	var userid int
	err = row.Scan(userid)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows found")
			return true
		} else {
			log.Fatal(err)

		}
	}
	return true
}

func (user *User) RetrieveToken(c *gin.Context) bool {
	log.Print("st 1")
	header := c.Request.Header.Get("token")
	if header == "" {
		log.Print("token empty")
		return false
	}
	tokenStr := header
	userid, _, err := accesstoken.ParseToken(tokenStr)
	if err != nil {
		fmt.Println("token denied")
		return false
	}
	id := int64(userid)
	print(id)
	// create user and check if user exists in db

	user.GetByUserId(id)
	if err != nil {

		fmt.Println("error in userid")
		return false

	}
	return true
}
func Updatetenant(token string, tenantid int) bool {
	stmt, err := database.Db.Prepare(updateenanttoken)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(token, tenantid)
	if err != nil {
		log.Fatal(err)
	}
	return true

}
func (u *User) Updateappuser() (bool, error) {
	stmt, err := database.Db.Prepare(updateappuser)
	if err != nil {
		fmt.Println(err)
		return false, err

	}

	_, err1 := stmt.Exec(&u.Email, &u.Mobile, &u.ID)
	if err1 != nil {
		fmt.Println(err1)
		return false, err1

	}
	return true, nil

}
func (u *User) Updateuserprofile() (bool, error) {
	stmt, err := database.Db.Prepare(updateuserprofile)
	if err != nil {
		fmt.Println(err)
		return false, err

	}

	_, err1 := stmt.Exec(&u.FirstName, &u.LastName, &u.Email, &u.Mobile, &u.Profileimage, &u.ID)
	if err1 != nil {
		fmt.Println(err1)
		return false, err1

	}
	return true, nil

}

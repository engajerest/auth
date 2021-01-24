package users

import (
	"github.com/engajerest/auth/logger"
	"github.com/engajerest/auth/utils/Errors"
	"github.com/engajerest/auth/utils/accesstoken"
	database "github.com/engajerest/auth/utils/dbconfig"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
) 


const(
	createUserQuery="INSERT INTO app_users (authname,password,hashsalt,contactno) VALUES(?,?,?,?)"
	insertUsertoProfileQuery="INSERT INTO app_userprofile (userid,firstname,lastname,email,contactno) VALUES(?,?,?,?,?)"
	getUseridByNameQuery="select user_id, email, mobile,status,created_date from engaje_users WHERE user_name = ?"
	authenticationQuery="select userid,password,hashsalt from app_users WHERE authname = ? and password = ?"
	usersGetAllQuery="select userid, firstname,lastname, contactno,email,status,created from app_userprofile"
	getUserByidQuery="select userid, firstname,lastname,contactno,email,status,created from app_userprofile WHERE userid=?"
	resetPasswordQuery="UPDATE app_users SET password=? ,hashsalt=?  WHERE userid = ?"
	insertTokentoSessionQuery="INSERT INTO app_session (userid,sessionname,sessiondate,sessionexpiry) VALUES(?,?,?,?)"
	checkUseridinSessionQuery="select userid from app_session WHERE userid= ?"
	userAuthentication="SELECT a.userid,b.firstname,b.lastname,b.email,b.contactno,b.status,b.created FROM app_users a, app_userprofile b WHERE a.userid=b.userid AND a.status ='Active' AND a.userid=?"

)



func (user *User) Create() int64 {
	fmt.Println("0")
	statement, err := database.Db.Prepare(createUserQuery)
	fmt.Println("1")
	if err != nil {
log.Fatal(err)
	}
	defer statement.Close()
	hashedPassword, err := HashPassword(user.Password)
	fmt.Println("2")

	res, err := statement.Exec(&user.Email,&user.Password,&hashedPassword,&user.Mobile)
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

func (user *User) InsertUserintoProfile() int64 {
	statement, err := database.Db.Prepare(insertUsertoProfileQuery)
	print(statement)
	
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	res, err := statement.Exec(&user.ID,&user.FirstName,&user.LastName,&user.Email,&user.Mobile)
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
	row := statement.QueryRow(user.FirstName, user.Password)
	print(row)

	var hashedPassword string

	err = row.Scan(&data.ID,&data.Password,&hashedPassword)
	print(err)
	fmt.Println("2")
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println(user)
	user.ID = data.ID
	return CheckPasswordHash(user.Password, hashedPassword)
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
		err := rows.Scan(&user.ID, &user.FirstName,&user.LastName, &user.Mobile, &user.Email, &user.Status,&user.CreatedDate)
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

func (user *User) GetByUserId(id int64) (*User,error) {
	
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
	return &data,nil
}
func  (user *User)  UserAuthentication(id int64) (*User,bool,error){
	fmt.Println("enrty in auth")
	print(id)
	var data User
	stmt, err := database.Db.Prepare(userAuthentication)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	err = row.Scan(&data.ID, &data.FirstName, &data.LastName, &data.Mobile, &data.Email, &data.Status, &data.CreatedDate)
	print(err)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows found")
		 data1:= Errors.RestError{}
		data1.Error=err
			return &data,false,err
		} else {
			log.Fatal(err)
			fmt.Println("nodata")
			var data1 *Errors.RestError
			data1.Error=err
			return &data,false,err
		}

	}
	// user.Check=true
	return &data,true,err
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
func (user *User)  InsertToken(token string) {
	// fmt.Println("0")
	statement, err := database.Db.Prepare(insertTokentoSessionQuery)
	// fmt.Println("1")
	defer statement.Close()
	if err != nil {
		log.Fatal(err)
	}
	sessiondate:=time.Now()
  sessionexpiry:=time.Hour.Hours()
	res, err := statement.Exec(&user.ID,token,sessiondate,sessionexpiry)
	if err != nil {
// print(token)
logger.Error("database error",err)
		log.Fatal(err)
	}
	_, err = res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("token inserted!")
	
}

func CheckStatusinSession(id int64) bool {
	stmt,err:=database.Db.Prepare(checkUseridinSessionQuery)
	if err != nil{
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
   var userid int
	err=row.Scan(userid)
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

func (user *User) RetrieveToken(c*gin.Context) bool{
	log.Print("st 1")
	header := c.Request.Header.Get("token")
if header == ""{
	log.Print("token empty")
	return false
}
      tokenStr := header
			userid, err := accesstoken.ParseToken(tokenStr)
			if err != nil {
			fmt.Println("token denied")
			return false
		}
		id:=int64(userid)
			   print(id)
			// create user and check if user exists in db
			
			user.GetByUserId(id)
			if err != nil {
				
				fmt.Println("error in userid")
				return false


			}
return true
}
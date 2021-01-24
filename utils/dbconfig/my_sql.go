package dbconfig



import(
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"

 )
 var Db *sql.DB

 func InitDB(dbname string,username string,password string,host string) {
	
	// Use root:dbpass@tcp(172.17.0.2)/hackernews, if you're using Windows.
	// db, err := sql.Open("mysql", "${dbName}:Package@123@tcp(198.24.174.194)/engajephoenix?parseTime=true")
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+")/"+dbname+"?parseTime=true")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
 		log.Panic(err)
	}
	Db = db
	log.Println("Database sucessfully configured")
 }

package database

import
(
	"errors"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"	

)

var Database *sql.DB

var Errors map[int64] string


func SetErrors() {

	Errors[1062] = "Email address already registered"

}

func GetError(e error) (error, string) {
	return errors.New("Something went wrong"),"dafook"
}


func Init() {
	db, err := sql.Open("mysql", "root@/dev.hyphen.io")
	if err != nil {

	}
	Database = db	
	log.Println("OH YEAH")
}
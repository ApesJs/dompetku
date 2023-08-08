package auth

import (
	"group-project/db_connection"

	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // go mod init dulu
)

type User struct {
	//id int primary key auto_increment not null,
	Username        string
	Password        string
	Email           string
	Fullname        string
	Balance         int
	Profile_picture string
	Address         string
	Phone_number    string
	Date_of_birth   string
}

func Register() {

	// KONEKSI DATABASE
	db, err := db_connection.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	newUser := User{}
	fmt.Print("Input Username:")
	fmt.Scanln(&newUser.Username) //jika mau menggunakan spasi menggunakan bufio
	fmt.Print("input Password:")
	fmt.Scanln(&newUser.Password)
	fmt.Print("input Email:")
	fmt.Scanln(&newUser.Email)
	fmt.Print("input Fullname:")
	fmt.Scanln(&newUser.Fullname)
	fmt.Print("input balance:")
	fmt.Scanln(&newUser.Balance)
	fmt.Print("input profile picture:")
	fmt.Scanln(&newUser.Profile_picture)
	fmt.Print("input Address:")
	fmt.Scanln(&newUser.Address)
	fmt.Print("input Phone Number:")
	fmt.Scanln(&newUser.Phone_number)
	fmt.Print("input date of birth:")
	fmt.Scanln(&newUser.Date_of_birth)

	statement, errPrepare := db.Prepare("INSERT INTO users (Username, Password, Email, Fullname, Balance, Profile_picture, Address, Phone_number, Date_of_birth) VALUES(?,?,?,?,?,?,?,?,?)")
	if errPrepare != nil {
		log.Fatal("error prepare insert", errPrepare.Error())
	}
	result, errInsert := statement.Exec(newUser.Username, newUser.Password, newUser.Email, newUser.Fullname, newUser.Balance, newUser.Profile_picture, newUser.Address, newUser.Phone_number, newUser.Date_of_birth)
	if errInsert != nil {
		log.Fatal("error insert data", errInsert.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("success insert data")
		} else {
			fmt.Println("failed to insert data")
		}
	}
}

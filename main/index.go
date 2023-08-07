package main

import (
	"fmt"
	"group-project/db_connection"
	"group-project/main/auth"
	"log"
	"os"
	"os/exec"
)

func main() {
	//KONEKSI DATABASE
	db, err := db_connection.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//MENU
	var menu int
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Print("Pilih Menu : ")
	fmt.Scan(&menu)

	//clear layar console
	clear := exec.Command("cmd", "/c", "cls")
	clear.Stdout = os.Stdout
	clear.Run()

	if menu == 1 {
		auth.Login()
	} else {
		fmt.Print("REGISTER")
	}
}

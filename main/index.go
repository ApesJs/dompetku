package main

import (
	"fmt"
	"group-project/helper"
	"group-project/main/auth"
	"time"
)

func main() {
	helper.ClearConsole()
	//MENU
	var menu int
	fmt.Println("")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("")
	fmt.Print("Pilih Menu : ")
	fmt.Scan(&menu)

	helper.ClearConsole()

	if menu == 1 {
		auth.Login()
	} else if menu == 2 {
		auth.Register()
	} else if menu == 0 {
		fmt.Println("Exit....")
		time.Sleep(2 * time.Second)
		helper.ClearConsole()
	} else {
		//REGISTER
	}
}

package main_menu

func MainMenu(channelMainMenu chan string) {
	// fmt.Println("")
	// fmt.Println("1. My Profile")
	// fmt.Println("2. Search Profile")
	// fmt.Println("3. Transaction")
	// fmt.Println("4. History Transaction")
	// fmt.Println("5. Logout")
	// fmt.Println("")
	// fmt.Println("0. Exit")
	// fmt.Println("")

	mainmenu := []string{
		"1. My Profile",
		"2. Search Profile",
		"3. Transaction",
		"4. History Transaction",
		"5. Logout",
		"",
		"0. Exit",
		"",
	}

	for i := 0; i < len(mainmenu); i++ {
		channelMainMenu <- mainmenu[i]
	}
	close(channelMainMenu)
}

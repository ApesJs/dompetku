package main_menu

func MainMenu(channelMainMenu chan string) {
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
	defer close(channelMainMenu)
}

package view

import (
	"fmt"
	"os"

	"github.com/KameiKojirou/popos-utils/core"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)


func BunMenu() {
	fmt.Println(lipgloss.NewStyle().Bold(true).Render("Bun Menu"))
	choice := ""
	form := huh.NewSelect[string]().
	Title("Select an option").
	Options(
		huh.NewOption("install", "Install Bun"),
		huh.NewOption("upgrade", "Upgrade Bun"),
		huh.NewOption("uninstall", "Uninstall Bun"),
		huh.NewOption("quit", "Quit"),
	).
	Value(&choice)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	switch choice {
		case "install":
			core.InstallBun()
		case "upgrade":
			// BunUpgrade()
		case "uninstall":
			// BunUninstall()
		case "q":
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
	}
}

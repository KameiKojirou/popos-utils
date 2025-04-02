package view

import (
	"fmt"
	"os"

	"github.com/KameiKojirou/popos-utils/core"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)	

func RustMenu() {
	fmt.Println(lipgloss.NewStyle().Bold(true).Render("Rust Menu"))
	choice := ""
	form := huh.NewSelect[string]().
	Title("Select an option").
	Options(
		huh.NewOption("Install Rust", "install"),
		huh.NewOption("Uninstall Rust", "uninstall"),
		huh.NewOption("Upgrade Rust", "upgrade"),
		huh.NewOption("Main Menu", "main"),
		huh.NewOption("Quit", "quit"),
	).
	Value(&choice)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	switch choice {
		case "install":
			core.InstallRust()
		case "uninstall":
			core.UninstallRust()
		case "upgrade":
			core.RustUpgrade()
		case "main":
			MainMenu()
		case "q":
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
	}
}

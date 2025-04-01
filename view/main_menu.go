package view

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func MainMenu() {
	fmt.Println(lipgloss.NewStyle().Bold(true).Render("Main Menu"))
	choice := ""
	form := huh.NewSelect[string]().
	Title("Select an option").
	Options(
		huh.NewOption("Bun", "bun"),
		huh.NewOption("Quit", "quit"),
	).
	Value(&choice)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	switch choice {
		case "bun":
			BunMenu()
		case "quit":
			log.Info(lipgloss.NewStyle().Bold(true).Render("Goodbye!"))
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
	}
}
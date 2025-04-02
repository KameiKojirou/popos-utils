/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/KameiKojirou/popos-utils/core"
	"github.com/KameiKojirou/popos-utils/view"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// rustCmd represents the rust command
var rustCmd = &cobra.Command{
	Use:   "rust",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
Run: func(cmd *cobra.Command, args []string) {
	log.Info(lipgloss.NewStyle().Bold(true).Render(fmt.Sprintf("Is Bun installed: %v", core.RustInstallCheck())))
	if len (args) == 0 {
		view.RustMenu()
	}
	if cmd.Flags().Changed("install") {
		core.InstallRust()
	} else if cmd.Flags().Changed("upgrade") {
		core.RustUpgrade()
	} else if cmd.Flags().Changed("uninstall") {
		core.UninstallRust()
	}
},
}

func init() {
	rootCmd.AddCommand(rustCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rustCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rustCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rustCmd.PersistentFlags().BoolP("install", "i", false, "Install Rust")
	rustCmd.PersistentFlags().BoolP("upgrade", "u", false, "Upgrade Rust")
}

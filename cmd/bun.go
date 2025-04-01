/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/KameiKojirou/popos-utils/core"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// bunCmd represents the bun command
var bunCmd = &cobra.Command{
	Use:   "bun",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info(lipgloss.NewStyle().Bold(true).Render(fmt.Sprintf("Is Bun installed: %v", core.BunInstallCheck())))
		if cmd.Flags().Changed("install") {
			log.Info("Installing Bun...")
		} else if cmd.Flags().Changed("upgrade") {
			log.Info("Upgrading Bun...")
		} else if cmd.Flags().Changed("uninstall") {
			log.Info("Uninstalling Bun...")
		}
	},
}

func init() {
	rootCmd.AddCommand(bunCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bunCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bunCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	bunCmd.PersistentFlags().BoolP("install", "i", false, "Install bun")
}

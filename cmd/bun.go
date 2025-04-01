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
	Short: "bun management",
	Long: `
		bun management commands to install, upgrade, and uninstall bun
		EXAMPLE USAGE:
		popos-utils bun --install
		popos-utils bun --upgrade
		popos-utils bun --uninstall
	`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info(lipgloss.NewStyle().Bold(true).Render(fmt.Sprintf("Is Bun installed: %v", core.BunInstallCheck())))
		if cmd.Flags().Changed("install") {
			core.InstallBun()
		} else if cmd.Flags().Changed("upgrade") {
			core.BunUpgrade()
		} else if cmd.Flags().Changed("uninstall") {
			core.BunUninstall()
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
	bunCmd.PersistentFlags().BoolP("upgrade", "u", false, "Upgrade bun")
}

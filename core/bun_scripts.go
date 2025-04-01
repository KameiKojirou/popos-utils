package core

import (
	"fmt"
	"os/exec"
	"github.com/charmbracelet/log"
)


func BunInstallCheck() bool {
	cmd := exec.Command("bun", "--version")
	err := cmd.Run()
	if err != nil {
		log.Error(err)
	}

	if cmd.ProcessState.ExitCode() != 0 {
		log.Error("Bun is not installed")
		return false
	}
	log.Info("Bun is installed")
	return true
}

func BunUpgrade() {
	cmd := exec.Command("bun", "upgrade")
	err := cmd.Run()
	if err != nil {
		log.Error(err)
	}
}


// install Bun
func InstallBun() {
	fmt.Println("Installing Bun...")
	cmd := exec.Command("curl", "https://bun.sh/install | bash")
	err := cmd.Run()
	if err != nil {
		log.Error(err)
	}
}
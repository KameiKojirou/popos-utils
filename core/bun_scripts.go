package core

import (
	"fmt"
	"os/exec"
	"os"
	"github.com/charmbracelet/log"
)


func BunInstallCheck() bool {
	cmd := exec.Command("bun", "--version")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
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
	fmt.Println("Upgrading Bun...")
	cmd := exec.Command("bun","upgrade")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

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


func BunUninstall() {
	fmt.Println("Uninstalling Bun...")
	cmd := exec.Command("rm", "-rf", "$HOME/.bun")
	err := cmd.Run()
	if err != nil {
		log.Error(err)
	}
}
package core

import (
	"fmt"
	"os/exec"
	"os"
	"github.com/charmbracelet/log"
)

func RustInstallCheck() bool {
	cmd := exec.Command("rustc", "--version")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Error(err)
	}
	if cmd.ProcessState.ExitCode() != 0 {
		log.Error("Rust is not installed")
		return false
	}
	log.Info("Rust is installed")
	return true
}

func InstallRust() {
	fmt.Println("Installing Rust...")
	cmd := exec.Command("curl", "--proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh")
	err := cmd.Run()
	if err != nil {
		log.Error(err)
	}
}

func UninstallRust() {
	fmt.Println("Uninstalling Rust...")
	// remove rust from .bashrc and .profile
	
	// remove rust
	cmd := exec.Command("rm", "-rf", "$HOME/.rustup")
	err := cmd.Run()
	if err != nil {
		log.Error(err)
	}
}

func RustUpgrade() {
	fmt.Println("Upgrading Rust...")
	cmd := exec.Command("rustup", "update")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Error(err)
	}
}
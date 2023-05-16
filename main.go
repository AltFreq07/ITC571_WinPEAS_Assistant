package main

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

//go:embed script.ps1 admin_wrapper.ps1
var script embed.FS

func main() {
	// Read the embedded scripts
	mainScriptData, err := script.ReadFile("script.ps1")
	if err != nil {
		fmt.Println("Error reading embedded script:", err)
		return
	}

	adminWrapperScriptData, err := script.ReadFile("admin_wrapper.ps1")
	if err != nil {
		fmt.Println("Error reading embedded admin_wrapper script:", err)
		return
	}

	// Get the directory of the executable
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}
	exeDir := filepath.Dir(exePath)

	// Write the embedded scripts to files in the same directory as the executable
	mainScriptPath := filepath.Join(exeDir, "temp_script.ps1")
	if err := os.WriteFile(mainScriptPath, mainScriptData, 0644); err != nil {
		fmt.Println("Error writing to main script file:", err)
		return
	}

	adminWrapperScriptPath := filepath.Join(exeDir, "check_admin.ps1")
	if err := os.WriteFile(adminWrapperScriptPath, adminWrapperScriptData, 0644); err != nil {
		fmt.Println("Error writing to admin wrapper script file:", err)
		return
	}

	// Run the admin_wrapper.ps1 script
	cmd := exec.Command("powershell.exe", "-ExecutionPolicy", "Bypass", "-File", adminWrapperScriptPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error executing PowerShell script:", err)
		return
	}

	// Remove the script files
	err = os.Remove(mainScriptPath)
	if err != nil {
		fmt.Println("Error removing main script file:", err)
		return
	}

	err = os.Remove(adminWrapperScriptPath)
	if err != nil {
		fmt.Println("Error removing admin wrapper script file:", err)
		return
	}
}

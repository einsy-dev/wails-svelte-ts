package shortcut

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func RegistryEdit() {
	// 1. Open the Registry Key
	key, _, err := registry.CreateKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Explorer\Advanced`, registry.SET_VALUE|registry.QUERY_VALUE)
	if err != nil {
		log.Fatalf("Error opening registry key: %v", err)
	}
	defer key.Close()

	// 2. Read existing disabled keys to avoid duplicates
	existing, _, _ := key.GetStringValue("DisabledHotkeys")
	if !strings.Contains(strings.ToUpper(existing), "V") {
		newVal := existing + "V"
		if err := key.SetStringValue("DisabledHotkeys", newVal); err != nil {
			log.Fatalf("Error writing to registry: %v", err)
		}
		fmt.Printf("Win+C added to DisabledHotkeys. New value: %s\n", newVal)
		restartExplorer()
	} else {
		fmt.Println("Win+C is already disabled in registry.")
	}
}

func restartExplorer() {
	fmt.Println("Restarting explorer.exe to apply changes...")
	// Kill explorer.exe
	exec.Command("taskkill", "/f", "/im", "explorer.exe").Run()
	// Start it again
	exec.Command("explorer.exe").Start()
}

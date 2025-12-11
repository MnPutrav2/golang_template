package cmd

import (
	"fmt"
	"os/exec"
)

func Fresh() {
	exec.Command("go", "run", "./cmd/rollback").Run()
	exec.Command("go", "run", "./cmd/migrate").Run()
	exec.Command("go", "run", "./cmd/seed").Run()

	fmt.Println("Migrate Fresh completed ===============")
}

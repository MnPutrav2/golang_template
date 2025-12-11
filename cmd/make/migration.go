package make

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Migration(name string) {
	if len(os.Args) < 2 {
		fmt.Println("Pake: go run make:migration NamaTable")
		return
	}

	// name := os.Args[1]
	slug := strings.ToLower(strings.ReplaceAll(name, " ", "_"))
	ts := time.Now().Format("20060102150405")

	base := fmt.Sprintf("%s_%s", ts, slug)
	up := base + ".up.sql"
	down := base + ".down.sql"

	dir := "db/migrations"
	os.MkdirAll(dir, 0o755)

	upPath := filepath.Join(dir, up)
	downPath := filepath.Join(dir, down)

	os.WriteFile(upPath, []byte("-- write up migration here\n"), 0o644)
	os.WriteFile(downPath, []byte("-- write down migration here\n"), 0o644)

	fmt.Println("Created:", upPath)
	fmt.Println("Created:", downPath)
}

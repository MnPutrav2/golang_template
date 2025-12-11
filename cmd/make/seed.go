package make

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Seeder(name string) {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run ./cmd/make/seed UserSeeder")
		return
	}

	// name := os.Args[1]
	slug := strings.ToLower(strings.ReplaceAll(name, " ", "_"))
	ts := time.Now().Format("20060102150405")

	file := fmt.Sprintf("%s_%s.sql", ts, slug)
	dir := "db/seed"
	os.MkdirAll(dir, 0o755)

	path := filepath.Join(dir, file)
	template := "-- seed SQL for " + name + "\n\n"
	os.WriteFile(path, []byte(template), 0o644)

	fmt.Println("Created seed:", path)
}

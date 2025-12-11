package main

import (
	"clean-arsitektur/cmd"
	"clean-arsitektur/cmd/make"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		Help()
		return
	}

	cli := os.Args[1]
	args := os.Args[2:]

	switch cli {
	case "server":
		cmd.Server()

	case "migrate":
		cmd.Migrate()

	case "rollback":
		cmd.Rollback()

	case "seed":
		cmd.Seed()

	case "fresh":
		cmd.Fresh()

	case "make:migration":
		if len(args) == 0 {
			fmt.Println("Usage: go run . make:migration <Name>")
			return
		}
		make.Migration(args[0])

	case "make:seeder":
		if len(args) == 0 {
			fmt.Println("Usage: go run . make:seeder <Name>")
			return
		}
		make.Seeder(args[0])

	case "make:template":
		if len(args) == 0 || len(args[1]) == 0 || len(args[2]) == 0 {
			fmt.Println("Usage: go run . make:template <folder_name> <file_name> <type>")
			return
		}

		make.Template(args[0], args[1], args[2])

	default:
		fmt.Print("Command not found")
	}
}

func Help() {
	fmt.Print(`
Available commands:
	go run . server							run server
	go run . migrate						run migration
	go run . rollback						down migration
	go run . seed							run seed SQL
	go run . fresh							rollback -> migrate -> seed
	go run . make:migration <table_name>				create timestamped up/down migration pair
	go run . make:seed <table_name>					create seed SQL template
	go run . make:template <create_folder> <file_name> <type>	create template [route, handler, service, repository]
	`)
}

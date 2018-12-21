package main

import (
	"flag"
	"os"

	"github.com/FreekingDean/bumper/database"
)

func main() {
	var path string
	flag.StringVar(&path, "database.path", "/var/lib/bumper/db.sqlite", "Path to bumper database")
	flag.Parse()

	db, err := database.NewDatabase(path)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	movies, err := movies.NewMovies(os.Getenv(""))
	if err != nil {
		panic(err)
	}
}

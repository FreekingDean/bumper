package main

import (
	"github.com/FreekingDean/bumper/database"
)

func main() {
	db, err := database.NewDatabase(path)
	defer db.Close()
	if err != nil {
		panic(err)
	}
}

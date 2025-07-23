package main

import "github.com/yourname/todoapp/internal/db"

func main() {
	db, err := db.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Your code here
}

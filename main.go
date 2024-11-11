package main

import (
	"20241111/database"
	"20241111/router"
	"log"
	"net/http"
)

func main() {
	db := database.DbOpen("20241111a")
	defer db.Close()
	r := router.NewRouter(db)

	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}

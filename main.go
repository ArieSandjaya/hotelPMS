package main

import (
	"hotelPMS/config"
	"hotelPMS/router"
)

func main() {
	db := config.Db()
	router := router.Router(db)

	router.Run(":8000")

}

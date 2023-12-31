package main

import (
	db "crud-mahasiswa/db"
	"crud-mahasiswa/routes"
)

func main() {
	db.InitDB()

	e := routes.InitRouter()

	e.Logger.Fatal(e.Start(":8080"))
}

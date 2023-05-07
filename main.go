package main

import (
	"alura-go-gin/database"
	"alura-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}

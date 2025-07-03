package main

import (
	"log"
	"net/http"
	"os"

	"laundry-branch/database"
	"laundry-branch/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	database.InitDB()
	r := router.SetupRouter()

	port := os.Getenv("PORT")
	log.Println("Servidor corriendo en puerto:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

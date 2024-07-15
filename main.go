package main

import (
	"fmt"
	"main/controllers"
	"main/db"
	"main/middlewares"
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db.Connect()

	routes()

	host := os.Getenv("HOST") + ":" + os.Getenv("PORT")

	ln, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening On " + host)
	if err = http.Serve(ln, nil); err != nil {
		panic(err)
	}
}

func routes() {

	http.HandleFunc("/", controllers.Redirect)
	http.HandleFunc("/ecpt", middlewares.Cors(controllers.Add))
}

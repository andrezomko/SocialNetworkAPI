package main

<<<<<<< HEAD
import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	fmt.Printf("Connected to database str: \n %s\n\n", config.ConnectionString)
	fmt.Println("API online!")

	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
=======
import "fmt"

func main() {
	fmt.Println("API running!")
}
>>>>>>> origin

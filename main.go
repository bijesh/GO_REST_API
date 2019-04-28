package main
 
import (
    "log"
	"net/http"
	"os"
	"GO_REST_API/route"
)
 
func main() {
	router := route.NewRouter()
	port := os.Getenv("PORT") //Get port from .env file
	if port == "" {
		port = "8080" //localhost
	}
    log.Fatal(http.ListenAndServe(":" + port, router))
 }


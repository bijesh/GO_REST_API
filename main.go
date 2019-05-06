package main
 
import (
    "log"
	"net/http"
	"os"
	"fmt"
	"GO_REST_API/route"
)
func recoverName() {  
    if r := recover(); r!= nil {
        fmt.Println("recovered from ", r)
    }
}
func main() {
	defer recoverName()
	router := route.NewRouter()
	port := os.Getenv("PORT") //Get port from .env file
	if port == "" {
		port = "8080" //localhost
	}
    log.Fatal(http.ListenAndServe(":" + port, router))
 }


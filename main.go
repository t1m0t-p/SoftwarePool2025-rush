package main

import (
    "fmt"
    "log"
    "net/http"
    "./router"
)

func main() {    
	r := router.Router()    
	fmt.Println("Starting server on the port 8080...")    
	http.ListenAndServe(":8080", r)
}
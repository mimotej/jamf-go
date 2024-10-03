package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World!")
}

func handleRequests(){
	http.HandleFunc("/", HelloWorld)
	fmt.Println("Serving on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


func main() {
	handleRequests()

}
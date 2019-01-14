package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// PORT of the server
const (
	PORT = 7082
)

// FileFolder reference
var (
	FileFolder string
	BaseURL    string
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Not enough arguments, path to file directory is expected")
	}

	FileFolder = os.Args[1]
	BaseURL = os.Args[2]

	log.SetOutput(os.Stdout)
	fmt.Println("Starting file server")
	r := initRoutes()
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), r)
}

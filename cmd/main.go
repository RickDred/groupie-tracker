package main

import (
	"fmt"
	"groupie-tracker/cmd/handlers"
	"net/http"
	"strconv"
)

const port = 8000

func main() {
	mux := http.NewServeMux()
	addr := ":" + strconv.Itoa(port)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/artist/", handlers.Artist)

	// starting server
	fmt.Printf("Server staring at localhost:%v\n", port)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println(err)
	}
}

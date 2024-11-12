package main

import (
	"fmt"
	"net/http"
	"os"

	Handlers "groupie/handlers"
)

const (
	PORT      = ":8080"
	Green     = "\033[1;32m "
	Red       = "\033[1;31m "
	ColorRest = "\033[0;0m"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", Handlers.HandleMain)
	http.HandleFunc("/artists", Handlers.HandleArtists)
	http.HandleFunc("/search", Handlers.HandleSearch)
	http.HandleFunc("/filter", Handlers.HandleFilter)

	// starting server
	fmt.Println(Green + "[+] Server started on http://localhost" + PORT + ColorRest)

	if err := http.ListenAndServe(PORT, nil); err != nil {
		fmt.Println(Red + "[-] Server error: " + err.Error() + ColorRest)
		os.Exit(2)
	}
}

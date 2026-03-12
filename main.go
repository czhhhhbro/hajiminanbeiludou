package main

import (

	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveIndex)
	http.Handle("/icon.png", http.StripPrefix("/icon.png", http.FileServer(http.Dir("."))))

	// Start server on port 8080
	log.Println("服务器启动在 http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// serveIndex handles serving the index.html file
func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

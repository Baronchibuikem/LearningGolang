package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// func serveDynamic(w http.ResponseWriter, r *http.Request) {
// 	response := "The time is now " + time.Now().string()
// 	fmt.Fprintln(w, response)
// }

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["id"]
	fileName := "files/" + pageID + ".html"
	_, err := os.Stat(fileName)
	if err != nil {
		fileName = "files/404.html"
	}
	http.ServeFile(w, r, fileName)
}

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/pages/{id:[0-9]+}", pageHandler)
	http.Handle("/", rtr)
	http.ListenAndServe(":8001", nil)
}

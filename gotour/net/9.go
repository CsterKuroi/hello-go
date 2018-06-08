package main

import (
	"log"
	"net/http"
)

func SayYeah(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Yeah"))
}

func main() {
	mux := http.NewServeMux()

	rh := http.RedirectHandler("http://www.baidu.com", 307)
	mux.Handle("/foo", rh)

	mux.HandleFunc("/bar", SayYeah)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}

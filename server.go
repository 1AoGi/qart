package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"qart/qrweb/web"
)

func main_old() {
	r := mux.NewRouter()
	r.HandleFunc("/qr", web.Draw)
	r.HandleFunc("/qr/show/{id}", web.Show)
	r.HandleFunc("/qr/draw", web.Draw)
	r.HandleFunc("/qr/arrow", web.Arrow)
	r.PathPrefix("/qr/static/").Handler(http.StripPrefix("/qr/static/", http.FileServer(http.Dir("./qrweb/public"))))
	http.Handle("/", r)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

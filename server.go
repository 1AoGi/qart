package main

import (
	"net/http"
	"qart/qrweb/web"
)

func main() {
	http.Handle("/qr", http.HandlerFunc(web.Draw))
	http.Handle("/qr/show/*", http.HandlerFunc(web.Show))
	http.Handle("/qr/draw", http.HandlerFunc(web.Draw))
	http.Handle("/qr/arrow", http.HandlerFunc(web.Arrow))
	http.Handle("/qr/static/", http.StripPrefix("/qr/static/", http.FileServer(http.Dir("./qrweb/public"))))
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

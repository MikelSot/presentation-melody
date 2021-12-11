package main

import "net/http"

//START OMIT
func (ro router) Router() {
	ro.webPage()
	ro.serveMux.HandleFunc("/ws", middleware.Log(ro.request)) // HL
}

func (ro router) request(w http.ResponseWriter, r *http.Request) {
	ro.handler.handler(ro.controller, w, r)
}

func (ro router) webPage() {
	ro.serveMux.Handle("/", http.FileServer(http.Dir("public"))) // HL
}

//END OMIT

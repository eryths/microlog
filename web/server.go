package web

import (
	"fmt"
	"microlog/web/actions"
	"net/http"
)

func Start() {

	http.HandleFunc("/", actions.Home)

	// static
	http.Handle("/static/", http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir("web/static"))))

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}

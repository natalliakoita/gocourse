package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type response struct {
	Cookie []string
}

func handlerCookie(w http.ResponseWriter, r *http.Request) {
	// Web server, responds to “/”
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// get value from Header[Cookie]
		var tm response
		x1 := r.Header["Cookie"]
		tm.Cookie = x1
		// load html witch data
		tmpl := template.Must(template.ParseFiles("form.html"))
		tmpl.Execute(w, tm)

	case http.MethodPost:
		// get data from form.html
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		name := r.FormValue("name")
		address := r.FormValue("address")
		// place time 2 minets
		expiration := time.Now().Add(120 * time.Second)
		// set cookie
		cookie := http.Cookie{Name: name, Value: address, Expires: expiration}
		http.SetCookie(w, &cookie)
		// redirect data to "/"
		http.Redirect(w, r, "http://localhost:8083/", http.StatusSeeOther)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	// declaration new endpoint
	http.HandleFunc("/", handlerCookie)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8083", nil); err != nil {
		log.Fatal(err)
	}
}

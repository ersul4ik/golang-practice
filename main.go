package main

import (
	"fmt"
	"net/http"
	"github.com/Pallinder/go-randomdata"
)


// HomeHandler return simple text
func (user *Customer) homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, user.str())
}

func main() {
	user := Customer{
		username: randomdata.SillyName(),
		age: 26,
		phoneNumber: randomdata.PhoneNumber(),
	}
	http.HandleFunc("/", user.homeHandler)
	http.ListenAndServe(":8000", nil)
}

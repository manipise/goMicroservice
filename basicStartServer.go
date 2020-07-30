package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello boss")
	})
	http.HandleFunc("/sayhi", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello ")

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Hello", string(b))
	})
	http.ListenAndServe(":8088", nil)
}

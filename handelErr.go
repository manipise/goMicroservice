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
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Header()
			// rw.Write([]byte("failed to get response"))
			// log.Println(err)
			http.Error(rw, "failed to get resoponse", http.StatusBadRequest)
			return
		}
		fmt.Println("Hello", string(b))
		fmt.Fprintf(rw, "Hello %s \n", b)
	})
	http.ListenAndServe(":8088", nil)
}

package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

		w.(http.Flusher).Flush()

		go func() {
			i := 0
			for {
				if i > 100 {
					break
				}
				fmt.Println("line:=", i)
				i++

				time.Sleep(time.Millisecond * 200)
			}
		}()

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		f, _ := w.(http.Flusher)

		for i := 0; i < 10; i++ {
			fmt.Fprintf(w, "time.now(): %v \n\r", time.Now())
			f.Flush()
			time.Sleep(time.Second)
		}

	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().UTC()
		stamp := ("The current machine timestamp in UTC: " + t.Format("2006-01-02T15:04:05.999999-07:00"))
		fmt.Println(stamp)
		fmt.Fprintf(w, stamp)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}

package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Eta(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query := r.URL.Query()
	fmt.Fprintf(w, "Your lat is %s and lon is %s\n", query.Get("lat"), query.Get("lon"))
}

func main() {
	router := httprouter.New()
	router.GET("/api/v1/cabs/eta", Eta)

	log.Fatal(http.ListenAndServe(":3000", router))
}

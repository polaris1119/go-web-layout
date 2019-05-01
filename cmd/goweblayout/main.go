package main

import (
	myhttp "github.com/polaris1119/go-web-layout/http"
	"log"
	"net/http"
)

func main() {
	myhttp.Route()

	log.Fatal(http.ListenAndServe(":6061", nil))
}

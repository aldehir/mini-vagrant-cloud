package main

import (
	"github.com/aldehir/mini-vagrant-cloud/server"
	"log"
)

func main() {
	s := server.NewBoxServer(":8080")
	log.Fatal(s.ListenAndServe())
}

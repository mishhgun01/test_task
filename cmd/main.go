package main

import (
	"github.com/gorilla/mux"
	"log"
	"test_task/pkg"
	"time"
)

func main() {
	s := pkg.NewStorage("127.0.0.1:6379", time.Hour)
	api := pkg.New(mux.NewRouter(), s)
	api.Handle()
	err := api.ListenAndServe("localhost:8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}

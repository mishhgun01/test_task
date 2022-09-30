package main

import (
	"github.com/gorilla/mux"
	"log"
	"test_task/pkg"
	"time"
)

func main() {
	s := pkg.NewStorage(CONN, time.Hour)
	defer s.Close()
	api := pkg.New(mux.NewRouter(), s)
	api.Handle()
	err := api.ListenAndServe(ADDR)
	if err != nil {
		log.Fatal(err.Error())
	}
}

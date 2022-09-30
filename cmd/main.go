package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"test_task/pkg"
	"time"
)

func main() {
	var conn string
	fmt.Print("enter connection:")
	_, err := fmt.Scanln(&conn)
	if err != nil {
		log.Fatal(err.Error())
	}
	s := pkg.NewStorage(conn, time.Hour*24)
	api := pkg.New(mux.NewRouter(), s)
	api.Handle()
	api.ListenAndServe("localhost:8080")
}

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"test_task/pkg"
	"time"
)

func main() {
	var conn, addr string
	fmt.Print("enter connection:")
	_, err := fmt.Scanln(&conn)
	fmt.Print("enter adress:port:")
	_, err = fmt.Scanln(&addr)
	if err != nil {
		log.Fatal(err.Error())
	}
	s := pkg.NewStorage(conn, time.Hour*24)
	api := pkg.New(mux.NewRouter(), s)
	api.Handle()
	err = api.ListenAndServe(addr)
	if err != nil {
		log.Fatal(err.Error())
	}
}

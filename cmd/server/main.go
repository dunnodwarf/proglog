package main

import (
	"fmt"
	"github.com/dunnodwarf/proglog/internal/server"
	"log"
)

func main() {
	fmt.Println("stop here")
	srv := server.NewHttpServer(":8080")
	log.Fatal(srv.ListenAndServe())
}

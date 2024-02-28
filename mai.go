package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("/Users/devops/tmp"))
	fmt.Println("subindo servidor na porta 9000...")
	log.Fatal(http.ListenAndServe(":9000", fs))
}

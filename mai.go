package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <diretorio> <porta>")
		os.Exit(1)
	}
	httpDir := os.Args[1]
	porta := os.Args[2]
	fs := http.FileServer(http.Dir(httpDir))
	fmt.Printf("Subindo servidor na porta %s ...", porta)
	log.Fatal(http.ListenAndServe(":"+porta, fs))
}

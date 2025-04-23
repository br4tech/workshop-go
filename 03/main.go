package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ola, mundo! Voce acesso a raiz do servidor.")
	})

	fmt.Println("Servidor inbiciado na porta 8080...")
	http.ListenAndServe(":8080", nil)
}

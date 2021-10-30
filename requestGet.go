package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Metodo ParseForm() gerou um erro: %v", err)
	}

	if (r.Method == "GET") {
		parametroEmailGet := r.Form.Get("email")
		fmt.Fprint(w, "Parametro email recebido por metodo GET.: " + string(parametroEmailGet))
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}


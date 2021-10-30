package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Metodo ParseForm() gerou um erro: %v", err)
	}

	if (r.Method == "POST") {
		parametroEmailPost := r.Form.Get("email")
		fmt.Fprint(w, "Parametro email recebido por metodo POST.: " + string(parametroEmailPost))
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}


package main

import (
  "fmt"
  "net/http"
  "github.com/lgdevlop/k8s-exemplo/go/message"
)

func imprimeMensagemBoasVindas(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, message.Greeting("Code.education Rocks!"))
}

func main() {
  fmt.Printf("Servidor escutando na porta 8000.\n")
  http.HandleFunc("/", imprimeMensagemBoasVindas)
  if err := http.ListenAndServe(":8000", nil); err != nil {
    panic(err)
  }
}

package message

import (
  "fmt"
)

func Greeting(mensagem string) string {
	return fmt.Sprintf("<b>%s</b>", mensagem)
}

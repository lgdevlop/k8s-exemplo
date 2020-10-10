package message

import (
	"testing"
)

func TestGreeting(testing *testing.T) {
	mensagemEsperada := "<b>Code.education Rocks!</b>"
	mensagem := Greeting("Code.education Rocks!")

	if mensagem != mensagemEsperada {
		testing.Errorf("Mensagem esperada é %v, mas o resultado foi %v", mensagemEsperada, mensagem)
	}
}

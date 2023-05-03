package documento

import (
	"strconv"
)

type Cpf struct {
	Identificacao Identificacao
}

func (c *Cpf) Generate() {
	// gerando 9 dígitos aleatórios
	r := GetRandomico()
	num1 := r.Intn(10)
	num2 := r.Intn(10)
	num3 := r.Intn(10)
	num4 := r.Intn(10)
	num5 := r.Intn(10)
	num6 := r.Intn(10)
	num7 := r.Intn(10)
	num8 := r.Intn(10)
	num9 := r.Intn(10)

	// calculando os dois dígitos verificadores
	digit1 := (num1*10 + num2*9 + num3*8 + num4*7 + num5*6 + num6*5 + num7*4 + num8*3 + num9*2) % 11
	if digit1 == 10 {
		digit1 = 0
	}

	digit2 := (num1*11 + num2*10 + num3*9 + num4*8 + num5*7 + num6*6 + num7*5 + num8*4 + num9*3 + digit1*2) % 11
	if digit2 == 10 {
		digit2 = 0
	}

	// formatando o CPF
	var cpf = strconv.Itoa(num1) + strconv.Itoa(num2) + strconv.Itoa(num3) + "." +
		strconv.Itoa(num4) + strconv.Itoa(num5) + strconv.Itoa(num6) + "." +
		strconv.Itoa(num7) + strconv.Itoa(num8) + strconv.Itoa(num9) + "-" +
		strconv.Itoa(digit1) + strconv.Itoa(digit2)

	c.Identificacao.Numero = cpf
}

func (c *Cpf) Validate() bool {
	return true
}

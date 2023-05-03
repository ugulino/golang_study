package documento

import (
	"strconv"
)

type Cnh struct {
	Identificacao Identificacao
}

func (c *Cnh) Generate() {
	r := GetRandomico()
	var cnh = ""
	for i := 0; i < 11; i++ {
		cnh += strconv.Itoa(r.Intn(10))
	}
	c.Identificacao.Numero = cnh
}

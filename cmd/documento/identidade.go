package documento

type Rg struct {
	Identificacao Identificacao
}

func (d *Rg) Generate() {
	r := GetRandomico()
	d.Identificacao.Numero = string(rune(r.Intn(900000) + 100000))
}

func (d *Rg) Validate() bool {
	return false
}

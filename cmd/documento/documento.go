package documento

type Identificacao struct {
	Numero   string
	Nome     string
	Emissao  string
	Validade string
}

type Operacoes interface {
	Generate()
	Validate() bool
}

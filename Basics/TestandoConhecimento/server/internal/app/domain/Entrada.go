package domain

type Entrada struct {
	IdEntrada       int
	EntNotaFiscal   int
	EntIdFornecedor int
	EntDataEntrada  string
	EntValorNota    float64
	EntChave        string
	EntIdNucleo     int
}

type DetEntrada struct {
	IdDetEntrada    int
	DenIdEntrada    int
	DenIdProduto    int
	DenQtd          int
	DenPrecoProduto float64
	DenIdNucleo     int
}

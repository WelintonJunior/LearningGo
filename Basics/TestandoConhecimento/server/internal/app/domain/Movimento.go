package domain

type Movimento struct {
	IdMovimento       int
	MovData           string
	MovIdVendedor     int
	MovSorteio        int
	MovHoraAbertura   string
	MovHoraFechamento string
	MovTotalVenda     float64
	MovIdUsuario      int
	MovEncerrado      string
	MovIdNucleo       int
}

type DetMovimento struct {
	IdDetalhe      int
	DetIdProduto   int
	DetIdMovimento int
	DetProQtd      float64
	DetHora        string
	DetProPreco    float64
}

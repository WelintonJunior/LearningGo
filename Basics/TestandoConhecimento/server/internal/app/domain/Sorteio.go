package domain

type Sorteio struct {
	IdSorteio   int
	SorData     string
	SorIdNucleo int
}

type DetSorteio struct {
	IdDetSorteio         int
	DetSorIdSorteio      int
	DetSorIdVendedor     int
	DetSorNumeroSorteado int
	DetSorPdvId          int
	DetSorMovEncerrado   string
	DetDorMovEnceHora    string
}

package domain

type Pdv struct {
	IdPDV             int
	PdvDesignacao     string
	PdvLocalFixo      string
	PdvCapacidade     int
	PdvSuspenso       int
	PdvIdNucleo       int
	PdvMotivoSuspenso string
}

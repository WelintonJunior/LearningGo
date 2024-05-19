package repository

import (
	db "example.com/teste/server/database"
	"example.com/teste/server/internal/app/domain"
)

type PdvRepository interface {
	PesquisarPdvPorId(idPdv, idNucleo, idUsuarioLogado, moduloPdv int64) (domain.Pdv, error)
}

type localPdvRepository struct{}

func NewLocalPdvRepository() *localPdvRepository {
	return &localPdvRepository{}
}

func (r *localPdvRepository) PesquisarPdvPorId(idPdv, idNucleo, idUsuarioLogado, moduloPdv int64) (domain.Pdv, error) {
	query := "SELECT * from tblPDV WHERE idPDV = ? and pdvIdNucleo = ?"
	row := db.DB.QueryRow(query, idPdv, idNucleo)

	var p domain.Pdv
	if err := row.Scan(&p.IdPDV, &p.PdvCapacidade, &p.PdvDesignacao, &p.PdvIdNucleo, &p.PdvLocalFixo, &p.PdvMotivoSuspenso, &p.PdvSuspenso); err != nil {
		return domain.Pdv{}, err
	}

	return p, nil
}

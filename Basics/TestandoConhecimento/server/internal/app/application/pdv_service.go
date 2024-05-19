package application

import (
	"example.com/teste/server/internal/adapteres/repository"
	"example.com/teste/server/internal/app/domain"
)

type PdvService struct {
	repo repository.PdvRepository
}

func NewPdvService(repo repository.PdvRepository) *PdvService {
	return &PdvService{repo: repo}
}

func (s *PdvService) PesquisarPdvPorId(idPdv, idNucleo, idUsuarioLogado, moduloPdv int64) (domain.Pdv, error) {
	return s.repo.PesquisarPdvPorId(idPdv, idNucleo, idUsuarioLogado, moduloPdv)
}

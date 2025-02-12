package service

import (
	"github.com/wazwki/api/internal/repository"
)

type NameService struct {
	repo repository.NameRepositoryInterface
}

func NewNameService(repo repository.NameRepositoryInterface) NameServiceInterface {
	return &NameService{repo: repo}
}

package service

import (
	"github.com/thiagoluiznunes/ze-challenge/domain/contract"
	"github.com/thiagoluiznunes/ze-challenge/infra/config"
)

// Service holds the domain service repositories
type Service struct {
	db  contract.DataManager
	cfg *config.Config
}

func New(db contract.DataManager, cfg *config.Config) (*Service, error) {

	svc := new(Service)
	svc.db = db
	svc.cfg = cfg

	return svc, nil
}

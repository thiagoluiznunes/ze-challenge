package data

import (
	"github.com/thiagoluiznunes/ze-challenge/data/datamongo"
	"github.com/thiagoluiznunes/ze-challenge/domain/contract"
	"github.com/thiagoluiznunes/ze-challenge/infra/config"
)

func Connect(cfg config.Config) (contract.DataManager, error) {
	return datamongo.Instance(cfg)
}

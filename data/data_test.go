package data

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluiznunes/ze-challenge/domain"
	"github.com/thiagoluiznunes/ze-challenge/infra/config"
)

func TestData(t *testing.T) {

	var cfg config.Config
	var err error

	err = json.Unmarshal([]byte(domain.MockConfig), &cfg)
	assert.Nil(t, err)
	db, err := Connect(cfg)
	assert.Nil(t, err)
	err = db.Close()
	assert.Nil(t, err)
}

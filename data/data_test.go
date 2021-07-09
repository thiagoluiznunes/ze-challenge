package data

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluiznunes/ze-challenge/infra/config"
)

func TestData(t *testing.T) {

	var cfg config.Config
	config := `{
		"app-name": "ze-delivery",
		"http-prefix": "/v1",
		"http-port": 5000,
		"db-host": "localhost",
		"db-port": 27017,
		"db-name": "ze_delivery",
		"db-user": "ze_user",
		"db-password": "hES6m2EXdjKqVkRf"
	}`
	err := json.Unmarshal([]byte(config), &cfg)
	assert.Nil(t, err)
	db, err := Connect(cfg)
	assert.Nil(t, err)
	err = db.Close()
	assert.Nil(t, err)
}

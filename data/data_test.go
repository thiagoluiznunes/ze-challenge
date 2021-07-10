package data

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluiznunes/ze-challenge/domain"
	"github.com/thiagoluiznunes/ze-challenge/infra/config"
)

func TestData(t *testing.T) {

	var cfg config.Config
	var err error
	defer func() {
		if recover() != nil {
			fmt.Println("fail creating connection", err)
		}
	}()

	config := domain.MockConfig
	err = json.Unmarshal([]byte(config), &cfg)
	assert.Nil(t, err)
	db, err := Connect(cfg)
	assert.Nil(t, err)
	err = db.Close()
	assert.Nil(t, err)
}

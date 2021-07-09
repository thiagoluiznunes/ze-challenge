package entity

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluiznunes/ze-challenge/domain"
)

func TestPartner(t *testing.T) {

	var partner Partner
	requestPartner := domain.MockRequestPartner
	err := json.Unmarshal([]byte(requestPartner), &partner)
	assert.Nil(t, err)
	err = partner.Validate()
	assert.Nil(t, err)
}

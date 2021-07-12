package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluiznunes/ze-challenge/data"
	"github.com/thiagoluiznunes/ze-challenge/domain"
	"github.com/thiagoluiznunes/ze-challenge/domain/contract"
	"github.com/thiagoluiznunes/ze-challenge/domain/entity"
	"github.com/thiagoluiznunes/ze-challenge/infra/config"
)

var err error
var srv *Service
var cfg config.Config
var db contract.DataManager

func TestService(t *testing.T) {
	// <setup code>
	t.Run("fail: test new service", TestNewService)
	t.Run("fail: test partner service", TestPartnerService)
	// <tear-down code>
}

func TestNewService(t *testing.T) {

	t.Run("fail: instance db", func(t *testing.T) {
		err = json.Unmarshal([]byte(domain.MockConfig), &cfg)
		assert.Nil(t, err)
		db, err = data.Connect(cfg)
		assert.Nil(t, err)
	})

	t.Run("fail: instance new service", func(t *testing.T) {
		srv, err = New(db, &cfg)
		assert.Nil(t, err)
		assert.NotEmpty(t, *srv)
	})
}
func TestPartnerService(t *testing.T) {

	var partner entity.Partner
	var partnerService *PartnerService

	t.Run("fail: instance new partner service", func(t *testing.T) {
		partnerService = NewPartnerService(srv)
		assert.NotEmpty(t, *partnerService)
	})

	t.Run("fail: instance partner entity", func(t *testing.T) {
		requestPartner := domain.MockRequestPartner
		err := json.Unmarshal([]byte(requestPartner), &partner)
		assert.Nil(t, err)
		err = partner.Validate()
		assert.Nil(t, err)
	})

	t.Run("fail: add new partner", func(t *testing.T) {
		seed := time.Now().UTC().UnixNano()
		partner.ID = fmt.Sprintf("test_datamongo_%d", seed)
		partner.Document = fmt.Sprintf("test_datamongo_document_%d", seed)
		_, err = partnerService.Add(context.TODO(), partner)
		assert.Nil(t, err)
	})

	t.Run("fail: add new partners in batch", func(t *testing.T) {
		partners := make([]entity.Partner, 2)
		err = json.Unmarshal([]byte(domain.MockRequestPartner), &partners[0])
		assert.Nil(t, err)
		err = json.Unmarshal([]byte(domain.MockRequestPartner), &partners[1])
		assert.Nil(t, err)
		assert.NotEmpty(t, partners)

		for index := range partners {
			seed := time.Now().UTC().UnixNano()
			partners[index].ID = fmt.Sprintf("test_datamongo_%d", seed)
			partners[index].Document = fmt.Sprintf("test_datamongo_document_%d", seed)
		}

		err = partnerService.AddInBatch(context.TODO(), partners)
		assert.Nil(t, err)
	})

	t.Run("fail: get partner by id", func(t *testing.T) {
		newParnter, err := partnerService.GetByID(context.TODO(), partner.ID)
		assert.Nil(t, err)
		assert.NotEmpty(t, newParnter)
		assert.Equal(t, partner, newParnter)
	})

	t.Run("fail: get all partners", func(t *testing.T) {
		allPartners, err := partnerService.GetAll(context.TODO())
		assert.Nil(t, err)
		assert.NotEmpty(t, allPartners)
	})
}

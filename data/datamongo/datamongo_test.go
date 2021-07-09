package datamongo

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/thiagoluiznunes/ze-challenge/domain"
	"github.com/thiagoluiznunes/ze-challenge/domain/contract"
	"github.com/thiagoluiznunes/ze-challenge/domain/entity"
	"github.com/thiagoluiznunes/ze-challenge/infra/config"
)

var err error
var cfg config.Config
var connManager contract.DataManager

func TestDataMongoRunner(t *testing.T) {
	// <setup code>
	t.Run("fail: test datamongo", TestDataMongo)
	t.Run("fail: test partner repo", TestPartnerRepo)
	// <tear-down code>
}

func TestDataMongo(t *testing.T) {

	t.Run("fail: instance config", func(t *testing.T) {
		config := domain.MockConfig
		err := json.Unmarshal([]byte(config), &cfg)
		assert.Nil(t, err)
	})

	t.Run("fail: instance datamanager", func(t *testing.T) {
		connManager, err = Instance(cfg)
		assert.Nil(t, err)
	})

	t.Run("fail: set indexes", func(t *testing.T) {
		err = connManager.SetIndexes()
		assert.Nil(t, err)
	})
}
func TestPartnerRepo(t *testing.T) {

	var partner entity.Partner
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
		err = connManager.Partner().Add(context.TODO(), partner)
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

		err = connManager.Partner().AddInBatch(context.TODO(), partners)
		assert.Nil(t, err)
	})

	t.Run("fail: get partner by id", func(t *testing.T) {
		newParnter, err := connManager.Partner().GetByID(context.TODO(), partner.ID)
		assert.Nil(t, err)
		assert.NotEmpty(t, newParnter)
		assert.Equal(t, partner, newParnter)
	})

	t.Run("fail: get all partners", func(t *testing.T) {
		allPartners, err := connManager.Partner().GetAll(context.TODO())
		assert.Nil(t, err)
		assert.NotEmpty(t, allPartners)
	})
}

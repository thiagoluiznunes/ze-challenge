package datamongo

import (
	"context"
	"sync"
	"time"

	"github.com/thiagoluiznunes/ze-challenge/domain/contract"
	"github.com/thiagoluiznunes/ze-challenge/infra/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	instance     *Conn
	onceDB       sync.Once
	onceInstance sync.Once
	connErr      error
)

// Conn is the MySQL connection manager
type Conn struct {
	client *mongo.Client
	db     *mongo.Database
	ctx    *context.Context
}

// Instance returns an instance of a DataManager
func Instance(cfg config.Config) (contract.DataManager, error) {

	onceInstance.Do(func() {
		conn, err := GetDB(cfg)
		if err != nil {
			connErr = err
			return
		}
		instance = &conn
	})

	return instance, connErr
}

func GetClientOptions(cfg config.Config) (clientOptions *options.ClientOptions) {

	uri := "mongodb://" + cfg.DBHost + ":" + cfg.DBPort
	credential := options.Credential{
		AuthSource: cfg.DBName,
		Username:   cfg.DBUser,
		Password:   cfg.DBPassword,
	}
	clientOptions = options.Client().ApplyURI(uri).SetAuth(credential)

	return clientOptions
}

func GetDB(cfg config.Config) (conn Conn, err error) {

	onceDB.Do(func() {
		clientOptions := GetClientOptions(cfg)
		conn.client, err = mongo.NewClient(clientOptions)
		if err != nil {
			connErr = err
			return
		}
		conn.db = conn.client.Database(cfg.DBName)

		context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err = conn.client.Connect(context)
		if err != nil {
			connErr = err
			return
		}

		err = conn.client.Ping(context, readpref.Primary())
		if err != nil {
			connErr = err
			return
		}

		conn.ctx = &context
	})

	return conn, connErr
}

func (c *Conn) SetIndexes() (err error) {

	err = setPartnerIndexes(c.db)
	if err != nil {
		return err
	}

	return nil
}

func (c *Conn) Close() (err error) {
	return c.client.Disconnect(*c.ctx)
}

func (c *Conn) Partner() contract.PartnerRepo {
	return newPartnerRepo(c.db)
}

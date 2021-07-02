package datamongo

import (
	"context"
	"sync"
	"time"

	"github.com/thiagoluiznunes/ze-challenge/domain/contract"
	"github.com/thiagoluiznunes/ze-challenge/infra/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	instance     *Conn
	dbInstance   *mongo.Client
	onceDB       sync.Once
	onceInstance sync.Once
	dbContext    *context.Context
	connErr      error
)

// Conn is the MySQL connection manager
type Conn struct {
	client *mongo.Client
	ctx    *context.Context
}

// Instance returns an instance of a DataManager
func Instance(cfg config.Config) (contract.DataManager, error) {

	onceInstance.Do(func() {
		client, context, err := GetDB(cfg)
		if err != nil {
			connErr = err
			return
		}
		instance = &Conn{client: client, ctx: context}
	})

	return instance, connErr
}

func GetClientOptions(cfg config.Config) (clientOptions *options.ClientOptions) {

	uri := "mongodb://" + cfg.DBHost + ":" + cfg.DBPort
	credential := options.Credential{
		Username: cfg.DBUser,
		Password: cfg.DBPassword,
	}
	clientOptions = options.Client().ApplyURI(uri).SetAuth(credential)

	return clientOptions
}

func GetDB(cfg config.Config) (client *mongo.Client, ctx *context.Context, err error) {

	onceDB.Do(func() {
		clientOptions := GetClientOptions(cfg)
		client, err = mongo.NewClient(clientOptions)
		if err != nil {
			connErr = err
			return
		}

		context, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(context)
		if err != nil {
			connErr = err
			return
		}

		dbInstance = client
		dbContext = &context
	})

	return dbInstance, dbContext, connErr
}

func (c *Conn) Close() (err error) {
	return c.client.Disconnect(*c.ctx)
}

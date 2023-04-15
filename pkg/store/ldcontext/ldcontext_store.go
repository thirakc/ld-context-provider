package ldcontext

import (
	"context"
	"ld-context-provider/pkg/logz"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var logger = logz.NewLogger()

type LDContextStore struct {
	*mongo.Client
	*mongo.Collection
}

func NewLDContextStore(uri, dbName, colName string) *LDContextStore {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		logger.Panic(err.Error())
	}

	coll := client.Database(dbName).Collection(colName)
	return &LDContextStore{client, coll}
}

func (s *LDContextStore) Save(docs any) (any, error) {
	res, err := s.InsertOne(context.TODO(), docs)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}

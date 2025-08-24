package tag

import (
  "context"
  "os"
  "log"
  "strings"
  "time"
  "errors"

  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

// @TODO: This file should go away

type DB struct {
  collection string
  connection *mongo.Client
}

func NewDb(collection string) *DB {
  return &DB{collection: collection}
}

func (db *DB) Connect() error {
  uri := os.Getenv("DB_CONNECTION_STRING")
  if !strings.HasPrefix(uri, "mongodb") {
    return errors.New("invalid connection string")
  }
  clientOptions := options.Client().ApplyURI(uri)

  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  var err error
  if db.connection, err = mongo.Connect(ctx, clientOptions); err != nil {
    return err
  }

  return db.connection.Ping(ctx, nil)
}

func (db *DB) Close() {
  db.connection.Disconnect(context.Background())
}

func (db *DB) Read(filter map[string]string, result *Tag) error {
  dataset := db.connection.Database("possohelp").Collection(db.collection)
  return dataset.FindOne(context.TODO(), filter).Decode(&result)
}


func (db *DB) AddValue(account, tagname, value string) error {
  filter := bson.M{"account": account, "name": tagname} 
  update := bson.M {
    "$push": bson.M {
      "values": TagValue{
        Value: value,
        Display: value,
        Matches: []string{value},
      },
    },
  }

  dataset := db.connection.Database("possohelp").Collection(db.collection)
  result, err := dataset.UpdateOne(context.TODO(), filter, update)
  if err != nil {
    return err
  }

  log.Printf("Matched %v document(s) and modified %v document(s).\n", result.MatchedCount, result.ModifiedCount)
  return nil
}

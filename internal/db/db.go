package db

import (
  "context"
  "log"
  "os"
  "strings"
  "time"

  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

var conn *mongo.Client
const DATABASE_NAME="possohelp"

func Connect() {
  log.Printf("connecting to database\n")
  uri := os.Getenv("DB_CONNECTION_STRING")
  if !strings.HasPrefix(uri, "mongodb") {
    log.Fatal("Invalid Connection String")
  }
  clientOptions := options.Client().ApplyURI(uri)
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  var err error
  conn, err = mongo.Connect(ctx, clientOptions)
  if err != nil {
    log.Fatal("Could not connect to DB: %v", err)
  }

  err = conn.Ping(ctx, nil)
  if err != nil {
    log.Fatalf("Could not ping DB: %v", err)
  }
  log.Printf("connected to database successfully\n")
}

func Disconnect() {
  log.Printf("disconnecting from database\n")
  conn.Disconnect(context.TODO())
  conn = nil
}

func GetCollection(collection string) *mongo.Collection {
  if conn == nil {
    Connect() 
  }
  log.Printf("returning collection: %v\n", collection)
	return conn.Database(DATABASE_NAME).Collection(collection)
}

func ReadOrdered(collection string, phoneNumbers []string) ([]bson.D, error) {
  dataset := GetCollection(collection)
  filter := bson.M{"phone": bson.M{"$in": phoneNumbers}}
  cursor, err := dataset.Find(context.Background(), filter)
  if err != nil {
    log.Fatal(err)
  }
  defer cursor.Close(context.Background())
  results := []bson.D{}
  if err = cursor.All(context.Background(), &results); err != nil {
    log.Fatal(err)
  }
  return results, nil
}

// No Order
func ReadUnordered(collection string, phoneNumbers []string) ([]bson.M, error) {
  dataset := GetCollection(collection)
  filter := bson.M{"phone": bson.M{"$in": phoneNumbers}}
  cursor, err := dataset.Find(context.Background(), filter)
  if err != nil {
    log.Fatal(err)
  }
  defer cursor.Close(context.Background())
  results := []bson.M{}
  if err = cursor.All(context.Background(), &results); err != nil {
    log.Fatal(err)
  }
  return results, nil
}

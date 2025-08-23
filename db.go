package main

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

var dbConn *mongo.Client
const DATABASE_NAME="possohelp"

func dbConnect() {
  println("dbConnect()")
  uri := os.Getenv("DB_CONNECTION_STRING")
  if !strings.HasPrefix(uri, "mongodb") {
    log.Fatal("Invalid Connection String")
  }
  clientOptions := options.Client().ApplyURI(uri)

  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  var err error
  dbConn, err = mongo.Connect(ctx, clientOptions)
  if err != nil {
    log.Printf("Could not connect to DB: %v", err)
  }

  err = dbConn.Ping(ctx, nil)
  if err != nil {
    log.Fatalf("Could not ping DB: %v", err)
  }
}

func dbDisconnect() {
  println("dbDisconnect()")
  dbConn.Disconnect(context.TODO())
  dbConn = nil
}

func dbGetCollection(collection string) *mongo.Collection {
  if dbConn == nil {
    dbConnect() 
  }
	return dbConn.Database(DATABASE_NAME).Collection(collection)
}

func dbReadOrdered(collection, phone string) ([]bson.D, error) {
  dataset := dbGetCollection(collection)
  filter := bson.M{"phone": phone}
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
func dbReadUnordered(collection, phone string) ([]bson.M, error) {
  dataset := dbGetCollection(collection)
  filter := bson.M{"phone": phone}
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

package db

import (
	"context"
	"errors"
	"os"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var credentials_file = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

//var client *firestore.Client = connect("./PRIVATE.json")
var client *firestore.Client = connect(credentials_file)

func connect(pathToCredentialsFile string) *firestore.Client {
	log.Debug("Database access credentials location: ", credentials_file)
	sa := option.WithCredentialsFile(pathToCredentialsFile)
	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Error("failed to create firebase client app")
		log.Fatal(err)
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Error("failed to create firestore client")
		log.Fatal(err)
	}
	return client
}

func disconnect() {
	defer client.Close()
}

func create(collection string, id string, doc map[string]interface{}) error {
	result, err := client.Collection(collection).Doc(id).Set(context.Background(), doc)
	if err != nil {
		log.Error("failed to create database document: ", doc, err)
		return err
	}
	log.Info("Successfully created database document: ", doc)
	log.Debug("firestore result response: ", result)
	return nil
}

func update(collection string, id string, doc map[string]interface{}) error {
	result, err := client.Collection(collection).Doc(id).Set(context.Background(), doc)
	if err != nil {
		return err
	}
	log.Info("Successfully updated database document: ", doc)
	log.Debug("firestore result response: ", result)
	return nil
}

func delete(collection string, id string) error {
	result, err := client.Collection(collection).Doc(id).Delete(context.Background())
	if err != nil {
		return err
	}
	log.Info("Successfully deleted database document: ", id)
	log.Debug("firestore result response: ", result)
	return nil
}

func get(collection string, id string) (map[string]interface{}, error) {
	result, err := client.Collection(collection).Doc(id).Get(context.Background())
	if err != nil {
		return nil, errors.New("failed to get document")
	}
	log.Info("Successfully got database document: ", id)
	log.Debug("firestore result response: ", result)
	return result.Data(), nil
}

func listAll(collection string) (*[]map[string]interface{}, error) {
	results := []map[string]interface{}{}
	iter := client.Collection(collection).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Error("failed to get document in list", err)
			continue
		}

		log.Debug("successfully added document to list: ", doc.Data())
		results = append(results, doc.Data())
	}
	log.Info("successfully got document list: ", results)
	return &results, nil
}

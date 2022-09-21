package devicedb

import (
	"context"
	"fmt"
	"log"
	"os"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var credentials_file = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

//var client *firestore.Client = connect("./PRIVATE.json")
var client *firestore.Client = connect(credentials_file)

func connect(pathToCredentialsFile string) *firestore.Client {
	fmt.Println("Credentials location:")
	fmt.Println(credentials_file)
	sa := option.WithCredentialsFile(pathToCredentialsFile)
	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatal(err)
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func Disconnect() {
	defer client.Close()
}

func Create(doc map[string]interface{}, collection string, id string) {
	result, err := client.Collection(collection).Doc(id).Set(context.Background(), doc)
	if err != nil {
		log.Fatalf("Failed creating document: %v Error: %v\n", doc, err)
	} else {
		fmt.Printf("Successfully created document: %v\n", doc)
	}
	fmt.Printf("Firestore result response: %v\n", result)
}

func Update(doc map[string]interface{}, collection string, id string) {
	result, err := client.Collection(collection).Doc(id).Set(context.Background(), doc)
	if err != nil {
		log.Fatalf("Failed updating document: %v Error: %v\n", doc, err)
	} else {
		fmt.Printf("Successfully updated document: %v\n", doc)
	}
	fmt.Printf("Firestore result response: %v\n", result)
}

func Delete(doc map[string]interface{}, collection string, id string) {
	result, err := client.Collection(collection).Doc(id).Delete(context.Background())
	if err != nil {
		log.Fatalf("Failed deleting document: %v Error: %v\n", doc, err)
	} else {
		fmt.Printf("Successfully deleted document: %v\n", doc)
	}
	fmt.Printf("Firestore result response: %v\n", result)
}

func Get(collection string, id string) map[string]interface{} {
	result, err := client.Collection(collection).Doc(id).Get(context.Background())
	if err != nil {
		log.Fatalf("Failed getting document. Error: %v\n", err)
	} else {
		fmt.Printf("Successfully got document!: %v\n", result.Data())
	}
	fmt.Printf("Firestore result response: %v\n", result)
	return result.Data()
}

func ListAll(collection string) []map[string]interface{} {
	results := []map[string]interface{}{}
	iter := client.Collection(collection).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error: %v\n", err)
			continue
		}

		fmt.Println("---")
		fmt.Println(doc.Data())
		results = append(results, doc.Data())
	}
	return results
}

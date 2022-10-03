package db

import (
	"encoding/json"

	"github.com/gregszalay/firestore-go/firego"

	"github.com/fatih/structs"
	log "github.com/sirupsen/logrus"
)

const COLLECTION string = "transactions"

func GetTransaction(id string) (*Transaction, error) {
	result, err := firego.Get(COLLECTION, id)

	jsonStr, err_marshal := json.Marshal(result)
	if err_marshal != nil {
		log.Error("failed to marshal transactionList", err_marshal)
	}

	var transaction Transaction
	if err_unmarshal := json.Unmarshal(jsonStr, &transaction); err != nil {
		log.Error("failed to unmarshal transaction json", err_unmarshal)

	}
	return &transaction, err
}

func ListTransactions() (*[]Transaction, error) {
	list, err := firego.ListAll(COLLECTION)
	transactionList := []Transaction{}
	for index, value := range *list {
		jsonStr, err := json.Marshal(value)
		if err != nil {
			log.Error("failed to marshal transactionList list element ", index, " error: ", err)
		}
		var tx Transaction
		if err := json.Unmarshal(jsonStr, &tx); err != nil {
			log.Error("failed to unmarshal transactionList list element ", index, " error: ", err)
		}
		transactionList = append(transactionList, tx)
	}
	log.Debug("List of transactions: ", transactionList)
	return &transactionList, err
}

func CreateTransaction(id string, newTransaction Transaction) error {
	return firego.Create(COLLECTION, id, structs.Map(newTransaction))
}

func UpdateTransaction(id string, newTransaction Transaction) error {
	return firego.Update(COLLECTION, id, structs.Map(newTransaction))
}

func DeleteTransaction(id string) error {
	return firego.Delete(COLLECTION, id)
}

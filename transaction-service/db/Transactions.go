package db

import (
	"encoding/json"

	"github.com/gregszalay/firestore-go/firego"

	log "github.com/sirupsen/logrus"
)

const COLLECTION string = "transactions"

func GetTransaction(id string) (*Transaction, error) {
	result, db_err := firego.Get(COLLECTION, id)
	if db_err != nil {
		log.Error("failed to get transaction from db", db_err)
		return nil, db_err
	}

	jsonStr, err_marshal := json.Marshal(result)
	if err_marshal != nil {
		log.Error("failed to marshal transaction", err_marshal)
		return nil, err_marshal
	}

	var transaction *Transaction = &Transaction{}
	if err_unmarshal := transaction.UnmarshalJSON(jsonStr); err_unmarshal != nil {
		log.Error("failed to unmarshal transaction json", err_unmarshal)
		return nil, err_unmarshal
	}

	return transaction, nil
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
	marshalled, marshal_err := json.Marshal(newTransaction)
	if marshal_err != nil {
		log.Error("CreateTransaction marshal error: ", marshal_err)
	}
	var unmarshalled map[string]interface{}
	unmarshal_err := json.Unmarshal(marshalled, &unmarshalled)
	if unmarshal_err != nil {
		log.Error("CreateTransaction unmarshal error: ", unmarshal_err)
	}
	return firego.Create(COLLECTION, id, unmarshalled)
}

func UpdateTransaction(id string, newTransaction Transaction) error {
	marshalled, marshal_err := json.Marshal(newTransaction)
	if marshal_err != nil {
		log.Error("CreateTransaction marshal error: ", marshal_err)
	}
	var unmarshalled map[string]interface{}
	unmarshal_err := json.Unmarshal(marshalled, &unmarshalled)
	if unmarshal_err != nil {
		log.Error("CreateTransaction unmarshal error: ", unmarshal_err)
	}
	return firego.Update(COLLECTION, id, unmarshalled)
}

func DeleteTransaction(id string) error {
	return firego.Delete(COLLECTION, id)
}

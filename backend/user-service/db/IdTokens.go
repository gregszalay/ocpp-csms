package db

import (
	"encoding/json"

	"github.com/gregszalay/firestore-go/firego"
	"github.com/gregszalay/ocpp-messages-go/types/AuthorizeRequest"

	log "github.com/sirupsen/logrus"
)

const COLLECTION string = "idTokens"

func GetIdToken(id string) (AuthorizeRequest.IdTokenType, error) {
	result, err := firego.Get(COLLECTION, id)

	jsonStr, err_marshal := json.Marshal(result)
	if err_marshal != nil {
		log.Error("failed to marshal IdTokenList", err_marshal)
	}

	var idTokenList AuthorizeRequest.IdTokenType
	if err_unmarshal := json.Unmarshal(jsonStr, &idTokenList); err != nil {
		log.Error("failed to unmarshal IdTokenList json", err_unmarshal)

	}
	return idTokenList, err
}

func ListIdTokens() (*[]AuthorizeRequest.IdTokenType, error) {
	list, err := firego.ListAll(COLLECTION)
	idTokenList := []AuthorizeRequest.IdTokenType{}
	for index, value := range *list {
		jsonStr, err := json.Marshal(value)
		if err != nil {
			log.Error("failed to marshal IdTokenList list element ", index, " error: ", err)
		}
		var idToken AuthorizeRequest.IdTokenType
		if err := json.Unmarshal(jsonStr, &idToken); err != nil {
			log.Error("failed to unmarshal IdTokenList list element ", index, " error: ", err)
		}
		idTokenList = append(idTokenList, idToken)
	}
	log.Debug("List of IdTokens: ", idTokenList)
	return &idTokenList, err
}

func CreateIdToken(id string, newIdToken AuthorizeRequest.IdTokenType) error {
	marshalled, marshal_err := json.Marshal(newIdToken)
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

func UpdateIdToken(id string, newIdToken AuthorizeRequest.IdTokenType) error {
	marshalled, marshal_err := json.Marshal(newIdToken)
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

func DeleteIdToken(id string) error {
	return firego.Delete(COLLECTION, id)
}

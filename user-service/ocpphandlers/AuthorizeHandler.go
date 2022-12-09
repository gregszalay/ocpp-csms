package ocpphandlers

import (
	"fmt"

	"github.com/gregszalay/ocpp-csms-common-types/QueuedMessage"
	"github.com/gregszalay/ocpp-csms/user-service/db"
	"github.com/gregszalay/ocpp-csms/user-service/publishing"
	"github.com/gregszalay/ocpp-messages-go/types/AuthorizeRequest"
	"github.com/gregszalay/ocpp-messages-go/types/AuthorizeResponse"
	"github.com/sanity-io/litter"
	log "github.com/sirupsen/logrus"
)

func AuthorizeHandler(request_json []byte, messageId string, deviceId string) {

	var req AuthorizeRequest.AuthorizeRequestJson
	payload_unmarshal_err := req.UnmarshalJSON(request_json)
	if payload_unmarshal_err != nil {
		fmt.Printf("Failed to unmarshal AuthorizeRequest message payload. Error: %s", payload_unmarshal_err)
	} else {
		fmt.Println("Payload as an OBJECT:")
		litter.Dump(req)
	}

	//TODO implement rfid auth properly
	all_tokens, err := db.ListIdTokens()
	if err != nil {
		log.Error("failed to get idtokens from db", err)
	}

	var status AuthorizeResponse.AuthorizationStatusEnumType_1 = AuthorizeResponse.AuthorizationStatusEnumType_1_Invalid
	for _, id_token := range *all_tokens {
		if id_token.IdToken == req.IdToken.IdToken {
			status = AuthorizeResponse.AuthorizationStatusEnumType_1_Accepted
		}
	}

	if status != AuthorizeResponse.AuthorizationStatusEnumType_1_Accepted{
		log.Warning("Authorization failed")
	}

	resp := AuthorizeResponse.AuthorizeResponseJson{
		IdTokenInfo: AuthorizeResponse.IdTokenInfoType{
			Status: status,
			EvseId: []int{0},
		},
	}
	qm := QueuedMessage.QueuedMessage{
		MessageId: messageId,
		DeviceId:  deviceId,
		Payload:   resp,
	}

	if err := publishing.Publish("AuthorizeResponse", qm); err != nil {
		log.Error("failed to publish AuthorizeResponse")
		log.Error(err)
	}

}

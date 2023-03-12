package events

import (
	"json"
	"log"
)

// Transfer object events occur you transfer an object from one address to another.
type TransferObject struct {
	PackageID         string
	TransactionModule string
	Sender            string
	Recipient         Recipient
	ObjectId          string
	Version           int
	Type              string
}

func (t TransferObject) unmarshal(raw TransferObject) TransferObject {

	err = json.Unmarshal(data, &raw)
	if err != nil {
		log.Println("Error unmarshalling json data:", err)
	}
	return
}

func (t TransferObject) decode()

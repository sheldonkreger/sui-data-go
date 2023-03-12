package events

// New object events occur for you create an object on the network.
type NewObject struct {
	PackageId         string
	TransactionModule string
	Sender            string
	Recipient         Recipient
	ObjectId          string
}

type Recipient struct {
	AddressOwner string
}

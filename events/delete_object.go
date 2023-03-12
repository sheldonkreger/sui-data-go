package events

// Delete object events occur when you delete an object.
type DeleteObject struct {
	PackageId         string
	TransactionModule string
	Sender            string
	ObjectId          string
}

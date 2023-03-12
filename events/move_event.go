package events

// Move calls emit Move events. You can define custom events in Move contracts.
type MoveEvent struct {
	PackageId         string
	TransactionModule string
	Sender            string
	Type              string
	Fields            Fields
	ObjectId          string
}

type Fields struct {
	Creator  string
	Name     string
	ObjectId string
}

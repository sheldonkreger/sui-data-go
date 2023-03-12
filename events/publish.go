package events

// Publish events occur when you publish a package to the network.
type Publish struct {
	Sender    string
	PackageId string
}

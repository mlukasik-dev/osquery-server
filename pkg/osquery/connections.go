package osquery

// ConnectionID is the type, which remote connection's IDs should be of.
type ConnectionID = string

type Connection struct {
	Client    *ExtensionManagerClient
	Connected bool
}

type Connections = map[ConnectionID]*Connection

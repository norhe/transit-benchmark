package vault

// Config : options to connect to Vault server
type Config struct {
	Address        string
	Token          string
	TransitKeyName string
}

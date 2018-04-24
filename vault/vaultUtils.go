package vault

import (
	"encoding/base64"
	"errors"
	"log"

	"github.com/norhe/transit-benchmark/workunit"

	"github.com/hashicorp/vault/api"
	"github.com/norhe/transit-benchmark/utils"
)

var vlt *api.Client
var keyName string

// reuse our client if we can
func initClient(vaultAddr, vaultToken string) *api.Client {
	if nil == vlt {
		cfg := api.DefaultConfig()
		cfg.Address = vaultAddr

		client, err := api.NewClient(cfg)
		utils.FailOnError(err, "Failed to initialize Vault client")

		client.SetToken(vaultToken)

		vlt = client
	}
}

// HandleCall : we want a generic way to invoke Vault calls.  The OperationType should
// tell the vault client what to do
func HandleCall(vaultAddr, vaultToken, keyName string, wu WorkUnit) (*WorkUnit, error) {
	initClient(vaultAddr, vaultToken)
	switch op := wu.OperationType; op {
	case workunit.OperationType.Encrypt:
		return encryptString(wu.Payload, keyName)
	case workunit.OperationType.SignData:
		return signData(wu.Payload, keyName)
	default:
		return nil, errors.New("no suitable OperationType found")
	}
}

/*func DecryptString(ciphertext string) ([]byte, error) {
	decrypted_contents, err := vlt.Logical().Write("transit/decrypt/"+keyName, map[string]interface{}{
		"ciphertext": ciphertext,
	})
	log.Printf("Decrypted: %+v", decrypted_contents)
	utils.FailOnError(err, "Error decrypting file: %s")

	decoded, err := base64.StdEncoding.DecodeString(decrypted_contents.Data["plaintext"].(string))
	utils.FailOnError(err, "Error decoding decrypted contents: %s")

	return decoded, err
}*/

func encryptString(ciphertext, keyName string) (string, error) {
	log.Printf("Encrypting: %s", ciphertext)

	// Payload must be base64 encoded before sending to Vault
	encoded := base64.StdEncoding.EncodeToString([]byte(ciphertext))

	log.Printf("Encoded: %s", encoded)

	// Write to Vault
	encryptedContents, err := vlt.Logical().Write("transit/encrypt/"+KEY_NAME, map[string]interface{}{
		"plaintext": encoded,
	})
	log.Printf("Encrypted: %+v", encrypted_contents)
	if err != nil {
		log.Fatalf("Error encrypting file: %s", err)
	}

	return encryptedContents.Data["ciphertext"].(string), err
}
